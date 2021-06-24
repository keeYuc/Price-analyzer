use serde::Deserialize;
use wasm_bindgen::prelude::*;
use yew::{
    format::{Json, Nothing},
    prelude::*,
    services::fetch::{FetchService, FetchTask, Request, Response},
};
#[derive(Deserialize, Debug, Clone)]
pub struct ISSPosition {
    latitude: String,
    longituede: String,
}
#[derive(Deserialize, Debug, Clone)]
pub struct ISS {
    message: String,
    timestamp: i32,
    iss_positon: ISSPosition,
}
#[derive(Debug)]
pub enum Msg {
    GetLocation,
    ReceiveResponse(Result<ISS, anyhow::Error>),
}

#[derive(Debug)]
pub struct Service {
    fetch_task: Option<FetchTask>,
    iss: Option<ISSPosition>,
    link: ComponentLink<Self>,
    error: Option<String>,
}

impl Service {
    fn view_iss_localtion(&self) -> Html {
        match self.iss {
            Some(ref space_station) => {
                html! {
                    <>
                    <p>{"message :"}</p>
                    <p>{format!("latittude{}",space_station.latitude)}</p>
                    <p>{format!("longitude{}",space_station.longituede)}</p>
                    </>
                }
            }
        }
    }
}
