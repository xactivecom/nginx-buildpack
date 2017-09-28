# Cloud Foundry NGINX Buildpack

This Cloud Foundry [buildpack](https://docs.cloudfoundry.org/buildpacks/) is a runtime version of NGINX which has been build using specific compile-time flags for serving HTTPS reverse proxy, or static file content.

### Buildpack User Documentation

Official buildpack documentation can be found at [staticfile buildpack docs](https://docs.cloudfoundry.org/buildpacks/staticfile/index.html).

### Building NGINX

The build of NGINX is performed on an Ubuntu 16 distribution.

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

1. Build and install NGINX

   Compile NGINX to /usr/local/sbin/nginx, and populate defult configuration to /usr/local/conf.

    ```bash
    make
    make install
    ```

1. Package NGINX for buildpack

   Copy the NGINX executible into the following directory structure:

    ```bash
    nginx/
    nginx/conf/
    nginx/logs/
    nginx/sbin/
    nginx/sbin/nginx
    sources.yml
    ```

   Create a compressed tarball, and compute MD5 hash:

    ```bash
    tar cvzf nginx-1.13.5-linux-x64.tgz nginx
    md5sum nginx-1.13.5-linux-x64.tgz
    ```

   Make the NGINX package available as an externally downloadable URL (details not shown). For example, https://myhealthrecords.services/dependencies/nginx/nginx-1.13.5-linux-x64.tgz.

1. Update buildpack manifest

   Update the buildpack dependencies nginx section of the manifest.yml with the NGINX package URL and MD5 hash.

    ```bash
    name: nginx
      version: 1.13.5
      uri: https://myhealthrecords.services/dependencies/nginx/nginx-1.13.5-linux-x64.tgz
      md5: b208e766191d113a142725a2b42f7a37
      ...
    ```

### Acknowledgements

This buildpack is based upon the Cloud Foundry [staticfile-buildpack](https://github.com/cloudfoundry/staticfile-buildpack).
