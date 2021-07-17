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
    iss: Option<ISS>,
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
                    <p>{format!("latittude{}",space_station.iss_positon.latitude)}</p>
                    <p>{format!("longitude{}",space_station.iss_positon.longituede)}</p>
                    </>
                }
            }
            None => {
                html! {
                    <button onclick=self.link.callback(|_|Msg::GetLocation)>
                    {"where is messafe"}
                    </button>
                }
            }
        }
    }
    fn view_fetching(&self) -> Html {
        if self.fetch_task.is_some() {
            html! {
                <p>{"fetching data ..."}</p>
            }
        } else {
            html! {
                 <p></p>
            }
        }
    }
    fn view_err(&self) -> Html {
        if let Some(ref error) = self.error {
            html! {
                <p>{error.clone()}</p>
            }
        } else {
            html! {}
        }
    }
}

impl Component for Service {
    type Message = Msg;
    type Properties = ();

    fn create(_props: Self::Properties, link: ComponentLink<Self>) -> Self {
        Self {
            fetch_task: None,
            iss: None,
            link,
            error: None,
        }
    }
    fn change(&mut self, _props: Self::Properties) -> bool {
        false
    }
    fn update(&mut self, msg: Self::Message) -> bool {
        use Msg::*;

        match msg {
            GetLocation => {
                // 1. build the request
                let request = Request::get("http://api.open-notify.org/iss-now.json")
                    .body(Nothing)
                    .expect("Could not build request.");
                // 2. construct a callback
                let callback =
                    self.link
                        .callback(|response: Response<Json<Result<ISS, anyhow::Error>>>| {
                            let Json(data) = response.into_body();
                            Msg::ReceiveResponse(data)
                        });
                // 3. pass the request and callback to the fetch service
                let task = FetchService::fetch(request, callback).expect("failed to start request");
                // 4. store the task so it isn't canceled immediately
                self.fetch_task = Some(task);
                // we want to redraw so that the page displays a 'fetching...' message to the user
                // so return 'true'
                true
            }
            ReceiveResponse(response) => {
                match response {
                    Ok(location) => {
                        self.iss = Some(location);
                    }
                    Err(error) => self.error = Some(error.to_string()),
                }
                self.fetch_task = None;
                // we want to redraw so that the page displays the location of the ISS instead of
                // 'fetching...'
                true
            }
        }
    }
    fn view(&self) -> Html {
        html! {
            <>
                { self.view_iss_localtion() }
                { self.view_fetching() }
                { self.view_err() }
            </>
        }
    }
}
