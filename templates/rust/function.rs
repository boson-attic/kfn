// kfn:dependency primal 0.2.3
// kfn:build-env RUSTC_WRAPPER=sccache
// kfn:build-dev true

use cloudevent::{Event, Reader, Writer};
use futures::IntoFuture;

// Function must have this signature
pub fn function(
    event: Option<Event>,
) -> Box<dyn futures::Future<Item = Option<Event>, Error = actix_web::Error>> {
    return match event.read_payload() {
        // A payload is here
        Some(Ok(p)) => Box::new(
            p.as_u64()
                // Payload is NaN
                .ok_or(actix_web::error::ErrorBadRequest("Expecting a number"))
                .map(|n| n as usize)
                // Calculate the nth prime
                .map(primal::StreamingSieve::nth_prime)
                // Serialize the calculated number
                .and_then(|n| {
                    serde_json::value::to_value(n as u64)
                        .map_err(actix_web::error::ErrorInternalServerError)
                })
                .and_then(|j| {
                    event
                        .unwrap()
                        // From the actual event, create a new one with new payload
                        .clone_with_new_payload("application/json", j)
                        .map(Option::from)
                        .map_err(actix_web::error::ErrorInternalServerError)
                })
                .into_future(),
        ),
        // Payload is malformed
        Some(Err(e)) => Box::new(futures::failed(actix_web::error::ErrorBadRequest(e))),
        // No payload at all
        None => Box::new(futures::failed(actix_web::error::ErrorBadRequest(
            "Expecting a non empty json payload",
        ))),
    };
}