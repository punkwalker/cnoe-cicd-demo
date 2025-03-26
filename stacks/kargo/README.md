# Kargo Stack

This directory contains a Jupyterhub deployment that's integrated with Keycloak

## Caveats
1) Reliance on `ref-implementation` for SSO
    - Another option is to use Password login. Follow the steps in [Kargo Documentation](https://docs.kargo.io/operator-guide/advanced-installation/advanced-with-argocd) for configuring admin user in Argo CD.
2) Reliance on cert-manager for Webhook certificates.
    - This can be removed with self-signed certificates. Check out [Github Issue](https://github.com/akuity/kargo/issues/2367) for more information.
3) First need to visit any component which starts with `

## Components
- Kargo
- Cert-Manager (for webhook certificates)

## Installation
Note: The stack is configured to use Keycloak for SSO; therefore, the ref-implementation is required for this to work.

```bash
idpbuilder create --use-path-routing  \
--package https://github.com/cnoe-io/stacks//ref-implementation \
--package https://github.com/cnoe-io/stacks//kargo
```

## Accessing Kargo
1) Kargo will be accessible at `https://kargo.cnoe.localtest.me:8443`.
2) Login using Keycloak SSO with `user1`. Get `user1` password using following command:
    ```
    idpbuilder get secrets -p keycloak -oyaml
    ```


