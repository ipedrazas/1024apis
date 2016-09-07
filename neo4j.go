package main

import (
	"encoding/csv"
	"os"
	"strings"
)

func toNeo(name string) string {
	return strings.Replace(name, "-", "", -1)
}

func writeServicesToCSV(services []Deployment, baseDir string) {
	file, err := os.Create(baseDir + "/result.csv")
	neo, err := os.Create(baseDir + "/neo4j.cypher")
	neoDep, err := os.Create(baseDir + "/neo4jDep.cypher")
	checkError("Cannot create file", err)
	defer file.Close()
	defer neo.Close()
	defer neoDep.Close()

	wr := csv.NewWriter(file)
	header := []string{"SERVICE", "DEPENDENCY"}
	wr.Write(header)
	for _, service := range services {
		neo.WriteString("CREATE (" + toNeo(service.Name) +
			":Service {name:'" + toNeo(service.Name) + "'})\n")
		if len(service.Dependencies) > 0 {
			neoDep.WriteString("CREATE\n")

			for idx, dep := range service.DependencyNames {
				neoDep.WriteString("\t(" + toNeo(service.Name) + ")-[:DEPENDS_ON]-> (" + toNeo(dep) + ")")

				if idx < (len(service.Dependencies) - 1) {
					neoDep.WriteString(",")
				}
				neoDep.WriteString("\n")

				record := []string{toNeo(service.Name), toNeo(dep)}
				wr.Write(record)
			}
		}

	}
	wr.Flush()
}
