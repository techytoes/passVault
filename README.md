### PassVault: Smart password manager.

#### Installation:
```
go get github.com/techytoes/passVault
```

#### Steps:

1. Run `init` command to create `creds.json` and `app.json` file in home directory.
2. Use `sniff` command to create new credentials.
3. Use `spit` command to fetch credentials.

#### Commands:

1. Initialize a new JSON file for storing credentials and a config file for storing app-level configs.
```
passvault init
```
Note: This command prompts to input a 32 byte string which would be used as an encryption key.
throughout the project.

2. Create a new credential using this command.

```
passvault sniff --app=<app/website>  --desc=<desc>
```

This command opens an interactive prompt where user can enter username and password.
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