# Versioning

To maintain a multi-module repository we need to have separate tagged versions of any submodule.

This can be achieved by adding tags that match the submodule like so:

```bash
git add <service-name>
git commit -am "Any changes that have been made"
git push
git tag <service-name>/<version-number>
git push origin <service-name>/<version-number>
```

Once this tag has been pushed the version that is included within the sibling module can be updated to reflect it.
