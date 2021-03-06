package audio

import "syscall/js"

// AudioContext represents an audio-processing graph
// built from audio modules linked together, each represented by an AudioNode.
// An audio context controls both the creation of the nodes it contains
// and the execution of the audio processing, or decoding.
// You need to create an AudioContext before you do anything else, as everything happens inside a context.
// https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext
// https://developer.mozilla.org/en-US/docs/Web/API/AudioContext
type AudioContext struct {
	js.Value
}

// GETTERS

// Current time returns an ever-increasing hardware time in seconds used for scheduling. It starts at 0.
// https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/currentTime
func (context AudioContext) CurrentTime() float64 {
	return context.Get("currentTime").Float()
}

// Destination is the final destination of all audio in the context.
// It often represents an actual audio-rendering device such as your device's speakers.
// https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/destination
func (context AudioContext) Destination() DestinationNode {
	node := AudioNode{Value: context.Get("destination")}
	return DestinationNode{AudioNode: node}
}

// SampleRate returns the sample rate (in samples per second) used by all nodes in this context.
// https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/sampleRate
func (context AudioContext) SampleRate() int {
	return context.Get("sampleRate").Int()
}

func (context AudioContext) State() AudioContextState {
	return AudioContextState(context.Get("state").String())
}

// METHODS

func (context AudioContext) Analyser() AnalyserNode {
	value := context.Call("createAnalyser")
	node := AudioNode{Value: value}
	return AnalyserNode{AudioNode: node}
}

func (context AudioContext) BiquadFilter() BiquadFilterNode {
	value := context.Call("createBiquadFilter")
	node := AudioNode{Value: value}
	return BiquadFilterNode{AudioNode: node}
}

// Gain creates a GainNode, which can be used to control the overall volume of the audio graph.
// https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/createGain
// https://developer.mozilla.org/en-US/docs/Web/API/GainNode
func (context AudioContext) Gain() GainNode {
	value := context.Call("createGain")
	node := AudioNode{Value: value}
	return GainNode{AudioNode: node}
}

// https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/createOscillator
// https://developer.mozilla.org/en-US/docs/Web/API/OscillatorNode
func (context AudioContext) Oscillator() OscillatorNode {
	value := context.Call("createOscillator")
	node := AudioNode{Value: value}
	return OscillatorNode{AudioNode: node}
}

// https://developer.mozilla.org/en-US/docs/Web/API/AudioContext/createMediaStreamSource
func (context AudioContext) MediaStreamSource(stream MediaStream) MediaStreamSourceNode {
	value := context.Call("createMediaStreamSource", stream.JSValue())
	node := AudioNode{Value: value}
	return MediaStreamSourceNode{AudioNode: node}
}

// https://developer.mozilla.org/en-US/docs/Web/API/AudioContext/resume
func (context AudioContext) Resume() {
	context.Call("resume")
}

// SUBTYPES

type AudioContextState string

const (
	AudioContextStateSuspended = AudioContextState("suspended")
	AudioContextStateRunning   = AudioContextState("running")
	AudioContextStateClosed    = AudioContextState("closed")
)
