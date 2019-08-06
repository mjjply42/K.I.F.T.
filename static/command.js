    let recording = document.getElementById("recording");

    startButton.addEventListener("click", function () {
      var audio = new Audio('record.wav');
      console.log("playing");
      audio.play();
      setTimeout(function(){  }, 200);
        $.ajax({
          url: "/command",
          method: "GET",
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
                  startSpotify();
                }
                else if (comm == "lights")
                {
                  checkLights(content);
                }
                else if (comm == "commands")
                {
                  $("#response").append('<li>' + "Command List:" +
                  "1.) Get me the weather---" +
                  "2.) Events near me---"+
                  "3.) Send email---"+
                  "4.) Set alarm---"+
                  "5.) Play music---"+
                  "6.) Turn on light---"+
                  "7.) Turn off light---"+
                  "8.) List commands---" + '</li>');
                  var voice_synth = window.speechSynthesis;
                  var voice_speech = new SpeechSynthesisUtterance("Here is the list of available commands for use");
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