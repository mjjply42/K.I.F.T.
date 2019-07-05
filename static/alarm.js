    let alarm_butt = document.querySelector(".alarm-start");     
    let alarm = document.querySelector(".bar-grow");             
    let left_side = document.querySelector(".left-side");         
    let right_side = document.querySelector(".right-side");
    let left_min = document.querySelector(".left-minute");
    let right_min = document.querySelector(".left-second");
    let left_sec = document.querySelector(".right-minute");
    let right_sec = document.querySelector(".right-second");

    $(".user-form").submit(function (e) {
      e.preventDefault();
    });
    function openAlarm(time) {
      console.log("pibidy");
      let add = 0;
      right_min.textContent = parseInt(time % 10);
      time /= 10;
      left_min.textContent= parseInt(time);
      let open = setInterval(animateWidth, 2);

      function animateWidth() {
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
      countdown();
      return ;
    }

    function  countdown()
    {
      console.log("pooidy");
        let alarmy = setInterval(runAlarm,1000);
        function  runAlarm()
        {
          if(right_sec.textContent == "0" && left_sec.textContent != 0)
          {
            right_sec.textContent = 9;
            left_sec.textContent -= 1;
          }
          else if(left_sec.textContent == "0" && right_sec.textContent == "0" && right_min.textContent != 0)
          {
            if(right_min.textContent > "0")
              right_min.textContent -= 1;
            left_sec.textContent = 5;
            right_sec.textContent = 9;
          }
          else if(right_min.textContent == "0" && right_sec.textContent == "0" && left_sec.textContent == "0" && right_min.textContent != 0)
          {
            if(left_min.textContent > "0")
              left_min.textContent -= 1;
            right_min.textContent = 9;
            left_sec.textContent = 5;
            right_sec.textContent = 9;
          }
          else if(left_min.textContent == "0" && right_min.textContent == "0" && right_sec.textContent == "0" && left_sec.textContent == 0)
          {
            clearInterval(alarmy);
            alarm.style.opacity = 1.0;
            let fade_out = setInterval(fader, 100);
            function  fader()
            {
              if(alarm.style.opacity > 0)
              {
                alarm.style.opacity -= .1;
              }
              else{
                clearInterval(fade_out);
                let state_set = document.querySelector(".alarm");
                state_set.id = "inactive";
                
                let item = document.querySelector(".bar-grow");
                item.parentNode.removeChild(item);

                let div = document.createElement("div");
                let div1 = document.createElement("div");
                let div2 = document.createElement("div");
                let h1 = document.createElement("h1");
                let h2 = document.createElement("h1");
                let h3 = document.createElement("h1");
                let h4 = document.createElement("h1");

                div.className = "bar-grow";
                div1.className = "left-side";
                div2.className = "right-side";
                h1.className = "left-minute";
                h2.className = "left-second";
                h3.className = "right-minute";
                h4.className = "right-second";

                document.querySelector(".clock").appendChild(div);
                document.querySelector(".bar-grow").appendChild(div1);
                document.querySelector(".bar-grow").appendChild(div2);
                document.querySelector(".left-side").appendChild(h1);
                document.querySelector(".left-side").appendChild(h2);
                document.querySelector(".right-side").appendChild(h3);
                document.querySelector(".right-side").appendChild(h4);
                
                return ;
              }
          }
        }
          else
          {
            right_sec.textContent -= 1;
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
      $(".client-input").focus();
      $(".client-input").on('keyup', function (event) {
        if (event.keyCode == 13) {
          user_input = parseInt($(".client-input").val());
          var item = document.querySelector(".client-input");
          item.parentNode.removeChild(item);
          var text = document.createElement("textarea");
          text.className = "client-input";
          text.style.display = "none";
          document.body.appendChild(text);
          $(".client-input").hide();
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
          openAlarm(user_input)
          return ;
        }
      });
    }