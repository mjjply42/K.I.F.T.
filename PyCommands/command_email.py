import smtplib
import sys

def email_sms(msg, dest, email, password):
    server = smtplib.SMTP('smtp.gmail.com', 587)
    server.starttls()
    server.login(email, password)

    server.sendmail(email, dest, msg)
    server.quit()


msg = sys.argv[1]
dest = sys.argv[2]
email = sys.argv[3]
password = sys.argv[4]

email_sms(msg, dest, email, password)