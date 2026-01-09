![Device Config](img/banner.png)

&nbsp;
&nbsp;
&nbsp;

#### Major Parameters

| Option | Type   | Description        |
| ------ | ------ | ------------------ |
| -i     | string | Device IP address  |
| -m     | string | Device MAC address |
| -p     | string | Device password    |

#### Read Requests

| Option | Type    | Description             |
| ------ | ------- | ----------------------- |
| -d     | boolean | Sow device details      |
| -fd    | boolean | Sow full device details |

#### Write Requests

| Option    | Type    | Description                                                                     |
| --------- | ------- | ------------------------------------------------------------------------------- |
| -rb       | boolean | Reboot device                                                                   |
| -rf       | boolean | Reset factory                                                                   |
| -rs       | boolean | Reset                                                                           |
| -nsp      | boolean | Set a new service password. The new password must be set using the option '-p'. |
| -csp      | string  | Change the service password                                                     |
| -clp      | string  | Change the live password                                                        |
| -cup      | string  | Change the user password                                                        |
| -cert     | boolean | Add certificate                                                                 |
| -skd      | string  | Set socket knocker destination                                                  |
| -cd       | string  | Set cloud destination                                                           |
| -cc       | string  | Cloud commission                                                                |
| -n        | string  | Set name                                                                        |
| -skm      | string  | Set socket knocker mode. Available modes: on, off, auto.                        |
| -sdt      | boolean | Sync Date/Time/UTC to PC                                                        |
| -vca      | string  | Set VCA profile                                                                 |
| -ro-on    | string  | Set Relay Output to On                                                          |
| -ro-off   | string  | Set Relay Output to Off                                                         |
| -ri-on    | string  | Set Relay Input to On                                                           |
| -ri-off   | string  | Set Relay Input to Off                                                          |
| -au-on    | string  | Set Audio to On                                                                 |
| -au-off   | string  | Set Audio to Off                                                                |
| -csrf-on  | string  | Set CSRF Protection to On                                                       |
| -csrf-off | string  | Set CSRF Protection to Off                                                      |

#### Preparing the device for a specific Remote Portal environment

| Option         | Type    | Description                    |
| -------------- | ------- | ------------------------------ |
| -rp-remoqa     | boolean | Prepare device to 'remoqa'     |
| -rp-remodevnew | boolean | Prepare device to 'remodevnew' |

#### Preparing the device for a specific Alarm Management environment

| Option       | Type    | Description                  |
| ------------ | ------- | ---------------------------- |
| -am-test     | boolean | Prepare device to 'test'     |
| -am-btuhtest | boolean | Prepare device to 'btuhtest' |
| -am-dev      | boolean | Prepare device to 'dev'      |
| -am-demo     | boolean | Prepare device to 'demo'     |

#### Output

| Option | Type    | Description                                     |
| ------ | ------- | ----------------------------------------------- |
| -so    | boolean | Simple output. No color, spinner or formatting. |
| -no    | boolean | No output                                       |
| -h     | boolean | Show help                                       |
| -v     | boolean | Show version                                    |
