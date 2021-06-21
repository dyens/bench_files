
#[macro_use] extern crate rocket;

use std::path::{Path};

use rocket::fs::{NamedFile};
 
#[get("/")]
pub async fn index() -> Option<NamedFile> {
    let path = Path::new("../my_file");
    NamedFile::open(path).await.ok()
}

#[launch]
fn rocket() -> _ {
    rocket::build().mount("/", routes![index])
}
