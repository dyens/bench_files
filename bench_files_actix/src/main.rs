
use actix_web::{web, App, HttpServer};
use actix_web::{HttpRequest, Result};
use actix_files::NamedFile;
use std::path::PathBuf;

async fn index(req: HttpRequest) -> Result<NamedFile> {
    let path: PathBuf = "../my_file".into();
    Ok(NamedFile::open(path)?)
}


#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
	    .route("/", web::get().to(index))
    })
    .bind("127.0.0.1:8080")?
    .run()
    .await
}
