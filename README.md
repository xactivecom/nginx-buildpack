# Cloud Foundry NGINX Buildpack

This Cloud Foundry [buildpack](https://docs.cloudfoundry.org/buildpacks/) is a runtime version of NGINX which has been build using specific compile-time flags for serving HTTPS reverse proxy, or static file content.

### Buildpack User Documentation

Official buildpack documentation can be found at [staticfile buildpack docs](https://docs.cloudfoundry.org/buildpacks/staticfile/index.html).

### Building NGINX

The build of NGINX is performed on an Ubuntu distribution.

1. Download and unpack the dependencies to the /usr/local directory

    ```bash
    cd /usr/local
    curl -OL https://ftp.pcre.org/pub/pcre/pcre-8.41.tar.gz
    tar xvzf pcre-8.41.tar.gz
    curl -OL http://www.zlib.net/zlib-1.2.11.tar.gz
    tar xvzf zlib-1.2.11.tar.gz
    curl -OL https://www.openssl.org/source/openssl-1.1.0f.tar.gz
    tar xvzf openssl-1.1.0f.tar.gz
    curl -OL http://nginx.org/download/nginx-1.13.5.tar.gz
    tar xvzf nginx-1.13.5.tar.gz
    ```

1. Configure the build for NGINX.

   ```bash
    cd nginx-1.13.5
   ./configure \
     --prefix=/usr/local \
     --with-http_ssl_module \
     --with-http_realip_module \
     --with-http_gunzip_module \
     --with-http_gzip_static_module \
     --with-http_auth_request_module \
     --with-http_random_index_module \
     --with-http_secure_link_module \
     --with-http_stub_status_module \
     --without-http_uwsgi_module \
     --without-http_scgi_module \
     --with-pcre-jit \
     --with-openssl=../openssl-1.1.0f \
     --with-pcre=../pcre-8.41 \
     --with-zlib=../zlib-1.2.11
   ```

1. Build NGINX

    ```bash
    make
    make install
    ```

1. Use in Cloud Foundry

   Upload the buildpack to your Cloud Foundry and optionally specify it by name

    ```bash
    cf create-buildpack [BUILDPACK_NAME] [BUILDPACK_ZIP_FILE_PATH] 1
    cf push my_app [-b BUILDPACK_NAME]
    ```

### Acknowledgements

This buildpack is based heavily upon Jordon Bedwell's Heroku buildpack and the modifications by David Laing for Cloud Foundry [nginx-buildpack](https://github.com/cloudfoundry-community/nginx-buildpack). It has been tuned for usability (configurable with `Staticfile`) and to be included as a default buildpack (detects `Staticfile` rather than the presence of an `index.html`). Thanks for the buildpack Jordon!
