package main

type worker struct {
}

func (w *worker) insertDataFromAzureIntoAWSCloud(clo cloud) {
	clo.insertIntoCloud()
}
