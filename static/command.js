    let recording = document.getElementById("recording");

    startButton.addEventListener("click", function () {
      var audio = new Audio('record.wav');
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