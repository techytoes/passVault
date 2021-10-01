### PassVault: Smart password manager.

#### Steps:

1. Run `init` command to create `creds.json` file.
2. Rename `sample-app.env` to `app.env` and set `ENC_KEY` variable. Remember to keep this string of length 32.
3. Use `sniff` command to create new credentials.
4. Use `spit` command to fetch credentials.

#### Commands:

1. Initialize a new JSON file for storing credentials.
```
passvault init
```

3. Create a new credential using this command.

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