package main

import (
	"cloud.google.com/go/firestore"
)

type db struct {
	firestoreClient *firestore.Client
}

func (d db) TrainingsCollection() *firestore.CollectionRef {
	return d.firestoreClient.Collection("trainings")
}
