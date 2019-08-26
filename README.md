# bosh-davcli

A CLI utility the BOSH Agent uses for accessing the [DAV blobstore](https://bosh.io/docs/director-configure-blobstore.html). 

Inside stemcells this binary is on the PATH as `bosh-blobstore-dav`.

### Developers

To update dependencies, use `gvt update`. Here is a typical invocation to update the `bosh-utils` dependency:

```
gvt update github.com/cloudfoundry/bosh-utils
```


# Pre-signed URLs

The command `sign` generates a pre-signed url for a specific object, action and duration:

`bosh-davcli <objectID> <action: get|put> <duration>`

The request will be signed using HMAC-SHA256 with a secret provided in configuration.

The HMAC format is:
```
{
    <HTTP Verb>\n
    <Object ID>\n
    <Unix timestamp of the signature time>\n
    <Unix timestamp of the expiration time>
}
```
The generated URL will be of format:

`https://some.url/object-id?st=the-HMAC-signature&ts=GenerationTimestamp&e=ExpirationTimestamp`
