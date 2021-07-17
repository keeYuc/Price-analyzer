use actix_web::{get, post, web, App, HttpResponse, HttpServer, Responder};

#[get("/")]
async fn hello() -> impl Responder {
    println!("get a request");
    HttpResponse::Ok().body("hello")
}

async fn test() -> impl Responder {
    println!("get a request");
    HttpResponse::Ok().body("test")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new().service(hello).route("/a", web::get().to(test)))
        .bind("127.0.0.1:8888")?
        .run()
        .await
}
