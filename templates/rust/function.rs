// kfn:dependency primal 0.2.3

use serde_json::json;

// Function must have this signature
pub fn function(event: Option<serde_json::Value>) -> impl futures::Future<Item=Option<serde_json::Value>, Error=actix_web::Error> {
    let input_json = event
        .unwrap_or(serde_json::Value::Null);
    let name = input_json
        .as_object()
        .and_then(|o| o.get("name"))
        .and_then(|v| v.as_str())
        .unwrap_or("World");
    let json = json!({
        "Hello": name,
        "A_fibonacci_number": primal::StreamingSieve::nth_prime(10)
    });
    futures::finished(Some(json))
}