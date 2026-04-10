# Dex website

* -- based on -- CNCF Hugo Starter
  * **[Hugo](https://gohugo.io/)**
    * allows
      * generating a static site
  * **[Docsy](https://www.docsy.dev/)**
    * provide
      * documentation theme
  * **[Netlify](https://www.netlify.com/)**
    * allows, about a DNS,
      * building
      * hosting

## documentation
* [here](layouts/index.md)

## how to run locally?

* TODO:
Make sure you have [npm](https://www.npmjs.com/) and [yarn](https://yarnpkg.com/) installed
* Clone this repository and run the following two commands in its directory:

```shell
# Run the server locally
make serve
```

## how to run | Netlify?

Netlify is a CI/CD build tool and hosting solution for (among other things) static sites
* We **strongly** recommend using Netlify unless you have a good reason not to.

This repository comes with a pre-configured [`netlify.toml`](https://github.com/cncf/hugo-netlify-starter/blob/master/netlify.toml) file
* To build to Netlify:

1. Go to [netlify.com](https://netlify.com) and sign up. We recommend signing up using a GitHub account.
2. Click **New Site from Git**, and give Netlify access to your GitHub account.
  > **Note:** For projects with lots of contributors, it can be handy to create a general/bot account instead of granting access with a personal account.

3. Install Netlify with access to your documentation site repository.
4. Leave all other settings as default and click **Deploy Site**.
