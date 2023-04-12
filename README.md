# Traefik IngressRoute Validate Webhook

This webhook use to check duplicate route in kubernetes cluster by admission webhook.

## Install

## Prerequisites

### Kubernetes API support

Kubernetes 1.9.0 or above with the `admissionregistration.k8s.io/v1` API enabled. Verify that by the following command:
```
kubectl api-versions | grep admissionregistration.k8s.io/v1
```
The result should be:
```
admissionregistration.k8s.io/v1
```

### Cert-Manager

This programs use cert-manager annotations to auto create cert.  
You can also use `Kubernetes CertificateSigningRequest` generate available cert.

## Install

Apply kubernetes files.
This project assume you install in `default` namespace,  
otherwise dnsName in `certificate.yaml` should be replaced with right namespace.

And you need to create a service to access traefik `traefik` port,  
it might not be exposed in default service.
In this case, a external service should be created.

Also, traefik default route must be applied,in that case you can access `/api` PathPrefix
with `traefik` port.
   
```
kubectl apply -k deployment
```

## How does it work?


Basically it just use Kubernetes Validation Webhook to check if `IngressRoute` could be applied
or not.

### Logical rules

1. Input: Traefik rule match just like this.

    ```
    Host(`foo.com`) && (PathPrefix(`/a`) || PathPrefix(`/b`))
    ```

2. Split and Combine: The splitAndCombine function is responsible for processing the input string and generating a list 
of minimal logical expressions. It does this by recursively splitting the input string into smaller segments and 
combining them based on the logical operators && (AND) and || (OR). For example, an input string like:  
    ```
    a && (b || c)
    ```
   will be transformed into:
    ```
    (a && b) || (a && c)
    ```
3. Parse Single Rule: The parseSingleRule function processes each logical expression in the list and extracts individual 
rules (such as `Host` and `PathPrefix`). It returns a map of rules where the keys are the rule functions, and the values 
are their corresponding arguments.
    ```
   map[string]string{"Host":"foo.com", "PathPrefix":"/a"}
    ```
   
    In that case, there is any difference between k and v of the map, it will be considered as a new route.
    For example, ``Host(\`foo.com\`) && PathPrefix(\`/v1.1.0\`)`` and ``Host(\`foo.com\`) && PathPrefix(\`/v1.0.0\`)`` are not the same

API is from traefik `/api/http/routers`