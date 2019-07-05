    let alarm_butt = document.querySelector(".alarm-start");     
    let alarm = document.querySelector(".bar-grow");             
    let left_side = document.querySelector(".left-side");         
    let right_side = document.querySelector(".right-side");       

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

    function  startAlarm(com_array)
    {
      $(".client-input").show();
      $("#response").append('<li>' + com_array + '</li>');
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
    }