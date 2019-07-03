let preview = document.getElementById("preview");
    let recording = document.getElementById("recording");
    let startButton = document.getElementById("startButton");
    let stopButton = document.getElementById("stopButton");
    let loginButton = document.querySelector("#login");
    let lastButton = document.querySelector("#last");
    let firstButton = document.querySelector("#first");
    let submit = document.querySelector("#submit");
    ////////////////////////////////////////////////////////////////
    let alarm_butt = document.querySelector(".alarm-start");      //
    let alarm = document.querySelector(".bar-grow");              //
    let left_side = document.querySelector(".left-side");         //  Alarm Buttons
    let right_side = document.querySelector(".right-side");       //
    ////////////////////////////////////////////////////////////////
    let spotify_butt = document.querySelector(".spotify-button");
    let recordingTimeMS = 4000;

    $(".user-form").submit(function (e) {
      e.preventDefault();
    });
    alarm_butt.onclick = function testAlarm() {
      let time = 32;
      let add = 0;
      console.log("damnn");
      let open = setInterval(animateWidth, 2);

      function animateWidth(time) {
        if (alarm.style.width != "120px") {
          alarm.style.width = add + 'px';
          add++;
        } else {
          clearInterval(open);
          add = 0;
          open = setInterval(animateHeight, 2);

          function animateHeight() {
            if (alarm.style.height != "80px") {
              alarm.style.height = add + 'px';
              add++;
            } else {
              clearInterval(open);
              let reveal = setInterval(showNumber, 30);
              add = 0;

              function showNumber() {
                if (left_side.style.opacity < 0.9 && right_side.style.opacity < 0.9) {
                  left_side.style.opacity = 0 + '.' + add;
                  right_side.style.opacity = 0 + '.' + add;
                  add++;
                } else {
                  clearInterval(reveal);
                }
              }
            }
          }
        }
      }

    }

    
    startButton.addEventListener("click", function () {
        $.ajax({
          url: "/command",
          method: "GET",
          success: function (data) {
            console.log(data)
            stringParse(data);

            function stringParse(command) {
              var com_array = command.split(";");
              var command = com_array[0];
              console.log(command);
              console.log(com_array);

              function dispatchComm(comm, content) {
                if (comm == "Alarm") {
                  $(".client-input").show();
                  $("#response").append('<li>' + com_array[1] + '</li>');
                  var alarm_synth = window.speechSynthesis;
                  var alarm_speech = new SpeechSynthesisUtterance("Enter Alarm Time");
                  alarm_speech.lang = 'en-US';
                  alarm_synth.speak(alarm_speech);
                  $(".client-input").on('keyup', function (event) {
                    if (event.keyCode == 13) {
                      user_input = parseInt($(".client-input").val());
                      $(".client-input").val("");
                      if (user_input > 59)
                        user_input = 59;
                      if (isNaN(user_input)) {
                        $("#response").append('<h1 class="error-mess" style="color: red;">' +
                          "ENTER A VALID NUMBER" + '</h1>');
                        setTimeout(function () {
                          let error_txt = document.querySelector("#response");
                          error_txt.removeChild(error_txt.lastChild);
                        }, 2000);
                      }
                      console.log(user_input);
                      //function openAlarm(user_input)
                      //{}
                    }
                  });
                } else if (comm == "Song") {
                  $(".spotify-button").show();
                  $("#response").append('<li>' + "Please Sign In First" + '</li>');
                  var song_synth = window.speechSynthesis;
                  var song_speech = new SpeechSynthesisUtterance("Sign In");
                  song_speech.lang = 'en-US';
                  song_synth.speak(song_speech);
                } else if (comm == "lights") {
                  if (content.substring(0,11).localeCompare(" Turning on") == 0) {
                    $('#image').attr("src", "https://images.cooltext.com/5285731.png")
                  } else {
                      $('#image').attr("src", "logo1.png")
                    }
                } else {
                  $("#response").append('<li>' + com_array[1] + '</li>');
                  var synth = window.speechSynthesis;
                  var test = document.querySelector("#response").lastChild.innerHTML;
                  var speech = new SpeechSynthesisUtterance(com_array[1]);
                  speech.lang = 'en-US';
                  synth.speak(speech);
                }
              }
              dispatchComm(command, com_array[1]);
            }
          },
        });
        })