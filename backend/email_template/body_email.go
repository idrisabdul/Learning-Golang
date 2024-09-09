package email_template

import (
	"time"
)

func BodyEmail(message string) string {
	indonesiaLocation, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().Local().In(indonesiaLocation)
	var body string = `
		<html lang="en" style="width: 100%; background-color: #f8f8fc">
            <body
              style="
                margin: 0 auto;
                font-family: Inter, sans-serif;
                background-color: #f8f8fc;
                font-size: 14px;
                font-style: normal;
                font-weight: 400;
                line-height: 20px;
                color: #667085;
                width: 100%;
              "
            >
              <div style="width: 100%">
                <table
                  style="
                    width: 70%;
                    justify-content: center;
                    margin: 0 auto;
                    border: 1px solid white;
                    border-collapse: collapse;
                  "
                >
                  <tbody style="background-color: #ffffff">
                    <tr>
                      <td style="text-align: center; vertical-align: middle" colspan="2">
                        <img
                          src="https://raw.githubusercontent.com/muhammadbahtiars/assets/main/Clip%20path%20group%20(1).png"
                          alt="Group"
                          style="max-width: 35%"
                        />
                      </td>
                    </tr>
                    <tr>
                      <td style="padding-left: 24px">
                        <p
                          style="
                            font-size: 20px;
                            font-style: normal;
                            font-weight: 400;
                            line-height: 20px;
                            color: #101828;
                          "
                        >
                          ` + message + `
                        </p>
                      </td>
                    </tr>
                    <tr>
                      <td
                        style="
                          padding-top: 50px;
                          background-color: #f8f8fc;
                          border: 1px solid #f8f8fc;
                          border-collapse: collapse;
                        "
                        colspan="2"
                      ></td>
                    </tr>
                    <tr>
                      <td style="text-align: center; vertical-align: middle" colspan="2">
                        <p>
                            ` + now.Format("2006") + `. PT Sigma Cipta Caraka - Telkomsigma. All Righs Reserved.
                        </p>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </body>
          </html>
	`
	return body
}
