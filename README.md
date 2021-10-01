### PassVault: Smart password manager.

#### Steps:

#### Commands:

1. Initialize a new JSON file for storing credentials.
```
passvault init
```

3. Create a new credential using this command.

```
passvault sniff --app=<app/website> --username=<username/email> --password=<password> --desc=<desc>
```
Note: `desc` is nullable field.

3. Print all usernames stored for the particular app.

```
passvault spit --app=<app/website>
```

4. Return password when a app and username combination is provided.

```
passvault spit --app=<app/website> --username=<username/email>
```

Note: instead of returning the password it is copied to the clipboard directly.