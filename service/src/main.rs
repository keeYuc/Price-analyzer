// extern crate redis;
use actix_web::{web, App, HttpRequest, HttpServer, Result};
use redis::Commands;
use serde::Deserialize;
use std::time::SystemTime;

const MAPKEY: &str = "status";

#[derive(Deserialize)]
struct Thing {
    thing_name: String,
}

async fn start(info: web::Json<Thing>) -> Result<String> {
    println!("get a request");
    match hget_map(&info.thing_name) {
        Ok(a) => {
            println!("yes value is");
            println!("{}", a)
        }
        Err(_) => match hset_map(&info.thing_name) {
            Ok(_) => println!("yes now to init"),
            Err(err) => println!("{}", format!("set err:{}", err)),
        },
    }
    Ok("Welcome start!".to_string())
}

async fn hello() -> Result<String> {
    println!("get a request");
    Ok("Welcome !".to_string())
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .route("/", web::get().to(hello))
            .route("/start", web::get().to(start))
    })
    .bind("127.0.0.1:8888")?
    .run()
    .await
}

fn hget_map(s: &String) -> redis::RedisResult<u32> {
    let client = redis::Client::open("redis://127.0.0.1/")?;
    let mut con = client.get_connection()?;
    con.hget(MAPKEY, s)
}

fn hset_map(s: &String) -> redis::RedisResult<u32> {
    let client = redis::Client::open("redis://127.0.0.1/")?;
    let mut con = client.get_connection()?;
    match SystemTime::now().duration_since(SystemTime::UNIX_EPOCH) {
        Ok(n) => con.hset(MAPKEY, s, n.as_secs()),
        Err(_) => panic!("SystemTime before UNIX EPOCH!"),
    }
}
