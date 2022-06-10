package main

type azureAdapter struct {
	azure *azure
}

func (ca *azureAdapter) insertIntoCloud() {
	ca.azure.insertIntoAzureCloud()
}
