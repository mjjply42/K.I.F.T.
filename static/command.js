    let recording = document.getElementById("recording");

    startButton.addEventListener("click", function () {
        $.ajax({
          url: "/command",
          method: "GET",
          success: function (data)
          {
            console.log(data)
            stringParse(data);

            function stringParse(command)
            {
              var com_array = command.split(";");
              var command = com_array[0];
              console.log("YO" + command);
              console.log(com_array);

              function dispatchComm(comm, content)
              {
                if (comm == "Alarm")
                {
                  stateCheck("alarm.js");
                  startAlarm(comm);
                } 
                else if (comm == "Song")
                {
                  //stateCheck("spotify.js");
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