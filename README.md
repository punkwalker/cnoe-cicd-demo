# ReInvent Kub308 CNOE

CNOE (Cloud Native Operational Excellence) is an open source organization comprised of enterrpise companies, sharing best practices on building, testing and deploying internal developer platforms (IDPs) at enterprise scale. 

This repository contains building blocks to deploy IDP capabilities, golden path templates/workflows and help you build your own Internal Developer Platform.

## Getting Started

### Install idpbuilder
To get started, you need to install [`idpbuilder`](https://github.com/cnoe-io/idpbuilder). See the [`instructions`] (https://github.com/cnoe-io/idpbuilder?tab=readme-ov-file#getting-started) in idpbuilder repo for details.

### Run idpbuilder

```bash
idpbuilder create --protocol http \
        --use-path-routing
```

idpbuilder creates local kind cluster with some core components to enable GitOps. These are the components installed,

* ArgoCD
* Gitea
* Ingress-Nginx


### Deploy stacks with building blocks to build IDP
```bash
idpbuilder create \
  --package reinvent-kub308-cnoe/stacks/cicd \
  --use-path-routing
```

The above command deploys IDP with below components.

* Argo Workflows
* Argo Events
* Backstage
* External Secrets
* Keycloak


## Security

See [CONTRIBUTING](CONTRIBUTING.md#security-issue-notifications) for more information.

## License

This project is licensed under the Apache-2.0 License.

