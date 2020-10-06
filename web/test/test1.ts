import { MongoClient, ObjectID, Collection } from "mongodb";

let clientPromise: Promise<MongoClient>;
export let client: MongoClient;

export async function connect(): Promise<void> {
    clientPromise = MongoClient.connect("mongodb://localhost/testDB", {
      useNewUrlParser: true,
      useUnifiedTopology: true,
    });
    client = await clientPromise;
    const db = client.db();
    const tasksCollection = db.collection("tasks");

    tasksCollection.insertOne({ device: 1, timestamp: 1 })
}

let a = connect();
