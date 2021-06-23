use std;
use yew::prelude::*;
use yew::services::fetch::Request;
pub struct App {
    link: ComponentLink<Self>,
    value: i64,
}

pub enum Msg {
    Increment,
    Decrement,
}

impl Component for App {
    type Message = Msg;
    type Properties = ();

    fn create(_props: Self::Properties, link: ComponentLink<Self>) -> Self {
        Self { link, value: 0 }
    }

    fn update(&mut self, msg: Self::Message) -> ShouldRender {
        match msg {
            Msg::Increment => {
                self.value += 1;
                true // Return true to cause the displayed change to update
            }
            Msg::Decrement => {
                self.value -= 1;
                true
            }
        }
    }

    fn change(&mut self, _props: Self::Properties) -> ShouldRender {
        false
    }

    fn view(&self) -> Html {
        html! {
            <div>
                <div class="panel">
                    // A button to send the Increment message
                    <button class="button" onclick=self.link.callback(|_| Msg::Increment)>
                        { "+1" }
                    </button>

                    // A button to send the Decrement message
                    <button onclick=self.link.callback(|_| Msg::Decrement)>
                        { "-1" }
                    </button>

                    // A button to send two Increment messages
                    <button onclick=self.link.batch_callback(|_| vec![Msg::Increment, Msg::Increment])>
                        { "+1, +1" }
                    </button>

                </div>

                // Display the current value of the counter
                <p class="counter">
                    { self.value }
                </p>

                // Display the current date and time the page was rendered
                <p class="footer">
                    { "Rendered: " }
                    {self.value}
                </p>
            </div>
        }
    }
}
