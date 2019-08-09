    let recording = document.getElementById("recording");

    startButton.addEventListener("click", function () {
      var audio = new Audio('record.wav');
      audio.play();
      setTimeout(function(){  }, 200);
        $.ajax({
          url: "/command",
          method: "GET",
          cache: false,
          success: function (data)
          {
            stringParse(data);

            function stringParse(command)
            {
              var com_array = command.split(";");
              var command = com_array[0];
              function dispatchComm(comm, content)
              {
                if (comm == "alarm")
                {
                  let state = stateCheck("alarm.js");
                  if(state == 1)
                    startAlarm(comm);
                  else
                    return ;
                } 
                else if (comm == "song")
                {
                  startSpotify(content);
                }
                else if (comm == "lights")
                {
                  checkLights(content);
                }
                else if (comm == "commands")
                {
                  $("#response").append('<li>' + "Command List:" + '<br>' + 
                  "1.) Get me the weather" + '<br>' +
                  "2.) Events near me"+ '<br>' +
                  "3.) Send email"+ '<br>' +
                  "4.) Set alarm"+ '<br>' +
                  "5.) Play music"+ '<br>' +
                  "6.) Turn on light"+ '<br>' +
                  "7.) Turn off light"+ '<br>' +
                  "8.) List commands" + '<br>' + '<br>' + '</li>');
                  var voice_synth = window.speechSynthesis;
                  var voice_speech = new SpeechSynthesisUtterance("Here is the list of available commands");
                  voice_speech.lang = 'en-US';
                  voice_synth.speak(voice_speech);
                }
                else
                {
                  sendResponse(comm, content);
                }
              }
              dispatchComm(command, com_array[1]);
            }
          }
        });
      });