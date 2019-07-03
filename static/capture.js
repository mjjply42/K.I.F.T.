let preview = document.getElementById("preview");
let recording = document.getElementById("recording");
let startButton = document.getElementById("startButton");
let stopButton = document.getElementById("stopButton");

const constraints = {audio: true, video: false};

let recordingTimeMS = 4000;

function wait(delayInMS) {
  return new Promise(resolve => setTimeout(resolve, delayInMS));
}
function startRecording(stream, lengthInMS) {
  let recorder = new MediaRecorder(stream);
  let data = [];
 
  recorder.ondataavailable = event => data.push(event.data);
  recorder.start();
 
  let stopped = new Promise((resolve, reject) => {
    recorder.onstop = resolve;
    recorder.onerror = event => reject(event.name);
  });

  let recorded = wait(lengthInMS).then(
    () => recorder.state == "recording" && recorder.stop()
  );
 
  return Promise.all([
    stopped,
    recorded
  ])
  .then(() => data);
}
function stop(stream) {
  stream.getTracks().forEach(track => track.stop());
}
startButton.addEventListener("click", function() {
  navigator.mediaDevices.getUserMedia(constraints).then(stream => {
    preview.srcObject = stream;
  }).then(() => startRecording(preview.captureStream(), recordingTimeMS))
  .then (recordedChunks => {
    let recordedBlob = new Blob(recordedChunks, { type: "audio/wav" });
    recording.src = URL.createObjectURL(recordedBlob);
    var a = document.createElement('a');
    document.body.appendChild(a);
    a.style = 'display:none';
    var url = window.URL.createObjectURL(recordedBlob);
    a.href = url;
    a.download = 'test.wav';
    a.click();
    window.URL.revokeObjectURL(url);
  })
}, false);
stopButton.addEventListener("click", function() {
  stop(preview.srcObject);
}, false);