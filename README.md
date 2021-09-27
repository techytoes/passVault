### PassVault: Smart password manager.


* How to store the data and what data to store.
  * website/app: can't be nullable.
  * username: nullable
  * email: nullable
  * password: not nullable
  * created: datetime(nullable)
  * last_used: datetime(nullable)
  * description: can be empty
* What commands do I provide.
  * create a username/pswd combination using cli. -> for now
    * `passvault sniff --application <app> --username <username> --password <pswd> --email <email> --desc <desc>`
    * `--password` is a compulsory flag.
    * either `--username` or `--email` is required to store the password.
    * `--desc` is optional field.
  * retrieve a username/pswd combination using cli. -> for now
    * `passvault spit --app <app>` -> shows all password for that app.
    * `passvault spit --app <app> --username <username>` -> shows password for that particular user.
    * `--app`: compulsory field.
  * bulk import from json. -> for later