import { connect } from "http2";
import MongoDB from "mongodb";

export let client: MongoDB.MongoClient;
function dbConn() {
    let conn = new MongoDB.MongoClient("mongodb://127.0.0.1:27017",{
    useNewUrlParser:true,
    useUnifiedTopology:true
    })

    conn.connect().then(v=>{
        return;
    }).catch(res=>{

    });
    

    if(conn.isConnected()){
        console.info("连接成功")

        let collection = conn.db("lnps").collection("auth")
        

    }else{
        console.info("连接失败")
    }
}

let con = dbConn()