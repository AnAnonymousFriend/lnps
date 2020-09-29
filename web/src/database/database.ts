import { connect } from "http2";
import Mongoose from "mongoose";

const url = "mongodb://localhost/testDB";

let database: Mongoose.Connection;

function Connect (){

    if (database) {
        return;
      }

    const db =  Mongoose.connect(url, {
        useNewUrlParser: true,
        useFindAndModify: true,
        useUnifiedTopology: true,
        useCreateIndex: true,
      });

      database = Mongoose.connection;
      database.once("open", async () => {
        console.log("Connected to database");
      });
      database.on("error", () => {
        console.log("Error connecting to database");
      });
}

let con = Connect()

