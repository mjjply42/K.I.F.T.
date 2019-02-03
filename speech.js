var synth = window.speechSynthesis;
var speech = new SpeechSynthesisUtterance("The weather is currently 59 degrees with a high of 75. Stay safe, folks!");
speech.lang = 'en-US';
let speak = document.querySelector("#speech");

speak.addEventListener('click', speakUp);

function speakUp()
{
    synth.speak(speech);
}

//Need to set 'SpeechSynthesisUtterance' to either a 
//text value that is pulled, or the actual 
//returned string
