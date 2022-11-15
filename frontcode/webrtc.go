package frontcode

import (
	"fmt"
	"io"
	"time"

	app "github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pion/webrtc/v3"

	"goappex/frontcode/signal"
)

type webrtcEx struct {
	app.Compo

	offerSessionDescription  string
	answerSessionDescription webrtc.SessionDescription
	renderReady              bool
	peerConnection           *webrtc.PeerConnection
	sendChannel              *webrtc.DataChannel
}

const messageSize = 15

// func (w *webrtcEx) _Render() app.UI {
// 	return app.Div().
// 		Class("container").
// 		Body(
// 			app.Div().
// 				Class("row").
// 				Body(
// 					app.H1().
// 						Body(
// 							app.Text("WebRTC demo"),
// 						),
// 				),
// 			app.Div().
// 				Class("row").
// 				Body(
// 					app.Div().
// 						Class("col-sm-6 no-left-padding").
// 						Body(
// 							app.H2().
// 								Body(
// 									app.Text("Offer"),
// 								),
// 							app.P().
// 								Body(
// 									app.I().
// 										Body(
// 											app.Text("Copy this to webrtc-cli stdin and press ^D"),
// 										),
// 								),
// 							app.Textarea().
// 								Class("form-control").
// 								ID("offer").
// 								Text(w.offerSessionDescription).
// 								ReadOnly(true),
// 						),
// 					app.Div().
// 						Class("col-sm-6 no-right-padding").
// 						Body(
// 							app.H2().
// 								Body(
// 									app.Text("Answer"),
// 								),
// 							app.P().
// 								Body(
// 									app.I().
// 										Body(
// 											app.Text("Paste webrtc-cli stdout here and press \"Start\""),
// 										),
// 								),
// 							app.Textarea().
// 								Class("form-control").
// 								ID("answer").
// 								OnChange(func(ctx app.Context, e app.Event) {
// 									fmt.Printf("answer text area changed: event %v\n", e.Value)
// 									v := ctx.JSSrc().Get("value").String()
// 									w.answerSessionDescription = v
// 									w.log(w.answerSessionDescription)
// 									// w.log("onchange!")
// 								}).
// 								OnInput(func(ctx app.Context, e app.Event) {
// 									// v := ctx.JSSrc().Get("value").String()
// 									// fmt.Printf("answer text area  on input : event %v\n", e)
// 									// fmt.Printf("answer text area  on input : value %s\n", v)
// 									// w.answerSessionDescription = v
// 									// w.log(w.answerSessionDescription)
// 									// w.log("oninput!")
// 								}),
// 						),
// 				),
// 			app.Div().
// 				Class("row").
// 				Body(
// 					app.H2().
// 						Body(
// 							app.Text("Session"),
// 						),
// 					app.Button().
// 						Class("btn btn-primary").
// 						Body(
// 							app.Text("Start"),
// 						),
// 					app.Button().
// 						Class("btn btn-primary").
// 						Body(
// 							app.Text("Stop"),
// 						),
// 					app.Button().
// 						Class("btn btn-primary").
// 						Body(
// 							app.Text("Refresh"),
// 						),
// 				),
// 			app.Div().
// 				Class("row").
// 				Body(
// 					app.H2().
// 						Body(
// 							app.Text("Tracks"),
// 						),
// 					app.Div().
// 						ID("tracks"),
// 				),
// 			app.Div().
// 				Class("row").
// 				Body(
// 					app.H2().
// 						Body(
// 							app.Text("Logs"),
// 						),
// 					app.Textarea().
// 						Class("form-control").
// 						ID("log").
// 						ReadOnly(true).
// 						OnFocus(func(ctx app.Context, e app.Event) {
// 							fmt.Println("in log on change...!!!")
// 							el := app.Window().GetElementByID("log")
// 							height := el.Get("scrollHeight")
// 							ofs := el.Get("offsetHeight")
// 							top := el.Get("scrollTop")
// 							fmt.Printf("height: %v\n", height)
// 							fmt.Printf("ofs: %v\n", ofs)
// 							fmt.Printf("top: %v\n", top)
// 							// el.Set("scrollTop"
// 						}),
// 				),
// 		)
// }

func (w *webrtcEx) Render() app.UI {
	fmt.Println("render invoked")
	return app.Div().
		Class("wrapper").
		Body(
			app.Text("Browser base64 Session Description"), app.Br(),

			app.Textarea().
				Class("form-control").
				ID("offerSessionDescription").
				// Text(w.offerSessionDescription).
				ReadOnly(true),
			app.Br(),
			app.Button().
				Class("btn btn-primary").
				Body(
					app.Text("Copy browser SDP to clipboard"),
				).
				OnClick(func(ctx app.Context, e app.Event) {
					fmt.Printf("button clicked event: %v\n", e)
					elem := app.Window().GetElementByID("offerSessionDescription")
					v := elem.Get("value")
					fmt.Println(v)
				}),
			// app.If(w.offerSessionDescription != "",
			// 	app.Textarea().
			// 		Class("form-control").
			// 		ID("localSessionDescriptionID").
			// 		Text(w.offerSessionDescription).
			// 		ReadOnly(true),
			// 	app.Br(),
			// 	app.Button().
			// 		Class("btn btn-primary").
			// 		Body(
			// 			app.Text("Copy browser SDP to clipboard"),
			// 		).
			// 		OnClick(func(ctx app.Context, e app.Event) {
			// 			fmt.Printf("button clicked event: %v\n", e)
			// 			elem := app.Window().GetElementByID("localSessionDescriptionID")
			// 			v := elem.Get("value")
			// 			fmt.Println(v)

			// 		}),
			// ).Else(
			// 	app.Text("local session description is not ready"),
			// ),
			app.Br(),
			app.Br(),
			app.Br(),
			app.Text("Golang base64 Session Description"),
			app.Br(),
			app.Textarea().
				Class("form-control").
				ID("remoteSessionDescription"),
			app.Br(),
			app.Button().
				Class("btn btn-primary").
				Body(
					app.Text("Start Session"),
				).OnClick(w.startSessionFunc),

			app.Br(),
			app.Br(),
			app.Text("Message"),
			app.Br(),
			app.Textarea().
				Class("form-control").
				ID("message").
				Body(
					app.Text("This is my DataChannel message!"),
				),
			app.Br(),
			app.Button().
				Class("btn btn-primary").
				Body(
					app.Text("Send Message"),
				).OnClick(w.sendMessageFunc),
			app.Br(),
			app.Br(),
			app.Text("Logs"),
			app.Br(),
			app.Div().Body(
				app.Textarea().
					Class("form-control").
					ID("logs").
					ReadOnly(true),
			),
		)
}

func newWebrtcEx() *webrtcEx {
	r := &webrtcEx{}
	return r
}

func (w *webrtcEx) OnInit() {
	fmt.Println("oninit invoked")

	var err error
	// Configure and create a new PeerConnection.
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
	w.peerConnection, err = webrtc.NewPeerConnection(config)
	if err != nil {
		w.handleError(err)
	}

	// Create DataChannel.
	w.sendChannel, err = w.peerConnection.CreateDataChannel("foo", nil)
	if err != nil {
		w.handleError(err)
	}
	w.sendChannel.OnClose(func() {
		fmt.Println("sendChannel has closed")
	})
	w.sendChannel.OnOpen(func() {
		fmt.Println("sendChannel has opened")

		candidatePair, err := w.peerConnection.SCTP().Transport().ICETransport().GetSelectedCandidatePair()

		fmt.Println(candidatePair)
		fmt.Println(err)
	})
	w.sendChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		w.log(fmt.Sprintf("Message from DataChannel %s payload %s", w.sendChannel.Label(), string(msg.Data)))
	})

	// Create offer
	offer, err := w.peerConnection.CreateOffer(nil)
	if err != nil {
		w.handleError(err)
	}
	if err := w.peerConnection.SetLocalDescription(offer); err != nil {
		w.handleError(err)
	}

	// Add handlers for setting up the connection.
	w.peerConnection.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		w.log(fmt.Sprint(state))
	})
	w.peerConnection.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate != nil {
			encodedDescr := signal.Encode(w.peerConnection.LocalDescription())
			el := app.Window().GetElementByID("offerSessionDescription")
			el.Set("value", encodedDescr)
		}
	})
}

// ReadLoop shows how to read from the datachannel directly
func (w *webrtcEx) ReadLoop(d io.Reader) {
	for {
		buffer := make([]byte, messageSize)
		n, err := d.Read(buffer)
		if err != nil {
			w.log(fmt.Sprintf("Datachannel closed; Exit the readloop: %v", err))
			return
		}

		w.log(fmt.Sprintf("Message from DataChannel: %s\n", string(buffer[:n])))
	}
}

// WriteLoop shows how to write to the datachannel directly
func (w *webrtcEx) WriteLoop(d io.Writer) {
	for range time.NewTicker(5 * time.Second).C {
		message := signal.RandSeq(messageSize)
		w.log(fmt.Sprintf("Sending %s \n", message))

		_, err := d.Write([]byte(message))
		if err != nil {
			w.handleError(err)
		}
	}
}

func (w *webrtcEx) log(msg string) {
	el := app.Window().GetElementByID("logs")
	// el.Set("innerHTML", el.Get("innerHTML").String()+msg+"<br>")
	el.Set("value", el.Get("value").String()+msg+"\n")

	height := el.Get("scrollHeight")
	ofs := el.Get("offsetHeight")
	// top := el.Get("scrollTop")
	// fmt.Printf("height: %v\n", height)
	// fmt.Printf("ofs: %v\n", ofs)
	// fmt.Printf("top: %v\n", top)

	// autoscroll to bottom
	el.Set("scrollTop", height.Int()-ofs.Int())

	// el.Set("scrollTop", top.Int())
	/*
		el := getElementByID("logs")
		el.Set("innerHTML", el.Get("innerHTML").String()+msg+"<br>")
	*/
}

func (w *webrtcEx) handleError(err error) {
	w.log("Unexpected error. Check console.")
	panic(err)
}

/*
	func getElementByID(id string) js.Value {
		return js.Global().Get("document").Call("getElementById", id)
	}
*/
func (w *webrtcEx) startSessionFunc(ctx app.Context, e app.Event) {
	el := app.Window().GetElementByID("remoteSessionDescription")
	sd := el.Get("value").String()
	if sd == "" {
		app.Window().Call("alert", "Session Description must not be empty")
		return
	}

	descr := webrtc.SessionDescription{}
	signal.Decode(sd, &descr)
	if err := w.peerConnection.SetRemoteDescription(descr); err != nil {
		w.handleError(err)
	}
}
func (w *webrtcEx) sendMessageFunc(ctx app.Context, e app.Event) {
	el := app.Window().GetElementByID("message")
	message := el.Get("value").String()
	if message == "" {
		app.Window().Call("alert", "Message must not be empty")
		return
	}
	if err := w.sendChannel.SendText(message); err != nil {
		w.handleError(err)
	}
}
