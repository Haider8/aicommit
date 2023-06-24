# aicommit
### Let AI write your commit messages!

![aicommit](https://github.com/Haider8/aicommit/assets/43299408/36052800-26e8-4f2a-9b15-49b2f063c294)



## Installation

1. Clone this repository on your local machine.
2. Install golang v1.20 and https://taskfile.dev for building binary.
3. Generate Chat GPT auth token from https://platform.openai.com/account/api-keys.
3. Now, execute these commands

    ```
    $ export API_KEY=<YOUR-AUTH-TOKEN>
    $ task build
    $ cp bin/aicommit /usr/local/bin
    ```
