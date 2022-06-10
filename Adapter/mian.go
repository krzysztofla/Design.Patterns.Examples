package main

func main() {
	worker := &worker{}
	aws := &aws{}
	azure := &azure{}

	worker.insertDataFromAzureIntoAWSCloud(aws)

	azureAdapterMachine := &azureAdapter{
		azure: azure,
	}

	azureAdapterMachine.azure.insertIntoAzureCloud()
}
