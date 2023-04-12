use actix_web::{web, App, HttpResponse, HttpServer, Responder};
use serde::Deserialize;

#[derive(Debug, Deserialize)]
pub struct Params {
    a: i32,
    b: i32,
}

async fn index(info: web::Query<Params>) -> impl Responder {
    let sum = info.a + info.b;
    println!("{}", sum);
    HttpResponse::Ok().body(format!("Сумма: {}", sum))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    println!("Сервер работает...");
    HttpServer::new(|| App::new().route("/", web::get().to(index)))
        .bind("127.0.0.1:8080")?
        .run()
        .await
}
