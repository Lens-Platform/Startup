# Security Policy
---

<p align="center">
  <strong>Reporting a vulnerability? See the <a href="#vulnerability-reporting">Vulnerability Reporting section</a></strong>
</p>

---

We understand that our users place a high level of trust in Blackspace. The security of Blackspace is a top priority.
That's why we apply widely accepted best practices when it comes to security.
This document will describe these practices and aims to be as transparent as
possible on our security efforts.

<!-- MarkdownTOC autolink="true" style="ordered" indent="   " -->

- [# Security Policy](#h1-idsecurity-policy-1security-policyh1)
- [Project Structure](#project-structure)
  - [Transparency](#transparency)
    - [Open Source](#open-source)
    - [Workflow](#workflow)
  - [Version Control](#version-control)
    - [Git](#git)
    - [Signed Commits](#signed-commits)
    - [Protected Branches](#protected-branches)
- [Personnel](#personnel)
  - [Education](#education)
  - [Policies](#policies)
  - [Two-factor Authentication](#two-factor-authentication)
- [Development](#development)
  - [Design & Architecture](#design--architecture)
    - [Golang](#golang)
    - [Unsafe Code](#unsafe-code)
    - [User Privileges](#user-privileges)
  - [Dependencies](#dependencies)
  - [Change Control](#change-control)
    - [Pull Requests](#pull-requests)
    - [Reviews & Approvals](#reviews--approvals)
    - [Merge Policies](#merge-policies)
    - [Automated Checks](#automated-checks)
      - [Vulnerability Scans](#vulnerability-scans)
      - [Fuzz Testing](#fuzz-testing)
- [Building & Releasing](#building--releasing)
  - [Network Security](#network-security)
  - [Runtime Isolation](#runtime-isolation)
  - [Asset Audit Logging](#asset-audit-logging)
  - [Asset Signatures & Checksums](#asset-signatures--checksums)
- [Vulnerability Reporting](#vulnerability-reporting)

<!-- /MarkdownTOC -->

## Project Structure

Project structure plays an important role in security. It creates guardrails
that prevent common security issues. This section will outline our deliberate
structural decisions that impact security.

### Transparency

We believe transparency is a strong deterrent of nefarious behavior that could
otherwise undermine security.

#### Open Source

Blackspace's dependencies are open-source. All code and changes are publicly
available at [our Github repo][urls.blackspace_repo].

#### Workflow

All of Blackspace's workflow is transparent.
[Pull requests][urls.blackspace_pull_requests], [issues][urls.blackspace_issues],
[chats][urls.blackspace_chat], and [our roadmap][urls.blackspace_roadmap]
are all publicly available.

### Version Control

Version control ensures that all code changes are audited and authentic.

#### Git

Blackspace leverages the [Git][urls.git] version-control system. This ensures all
changes are audited and traceable.

#### Signed Commits

Because of Blackspace's [merge style](CONTRIBUTING.md#merge-style), commits to
release branches are signed by Github itself during the squash and merge
process. Commits to development branches are encouraged to be signed but not
required since changes must go through a [review process](#reviews--approvals).

#### Protected Branches

Blackspace cuts releases from the `master` and `v*` branches _only_. These branches
are protected. The exact requirements are:

* Cannot be deleted.
* Force pushes are not allowed.
* A linear history is required.
* Signed commits are required.
* Administrators are included in these checks.

## Personnel

### Education

Blackspace team members are required to review this security document as well as
the [contributing](CONTRIBUTING.md) and [reviewing](REVIEWING.md) documents.

### Policies

Blackspace maintains this security policy. Changes are communicated to all Blackspace
team members.

### Two-factor Authentication

All Blackspace team members are required to enable two-factor authentication
for their Github accounts.

## Development

### Design & Architecture

The base of Blackspace's security lies in our choice of underlying technology and
decisions around design and architecture.

#### Golang

The [Golang programming language][urls.go] is not memory and thread-safe; it will
not catch many common sources of concurency vulnerabilities at compile time. Hence, it is the responsibility of the developer
to properly account for any potential race conditions.

#### Unsafe Code

Blackspace does not allow the use of unsafe code except in circumstances where it
is required, such as dealing with CFFI.

#### User Privileges

Blackspace is always designed to run under non-`root` privileges, and our
documentation always defaults to non-`root` use.

### Dependencies

Blackspace aims to reduce the number of dependencies it relies on. If a dependency
is added it goes through a comprehensive review process that is detailed in
the [Reviewing guide](REVIEWING.md#dependencies).

### Change Control

As noted above Blackspace uses the Git version control system on Github.

#### Pull Requests

All changes to Blackspace must go through a pull request review process.

#### Reviews & Approvals

All pull requests must be reviewed by at least one Blackspace team member. The
review process takes into account many factors, all of which are detailed in
our [Reviewing guide](REVIEWING.md). In exceptional circumstances, this
approval can be retroactive.

#### Merge Policies

Blackspace requires pull requests to pass all [automated checks](#automated-checks).
Once passed, the pull request must be squashed and merged. This creates a clean
linear history with a Blackspace team member's co-sign.

#### Automated Checks

When possible, we'll create automated checks to enforce security policies.

##### Vulnerability Scans

Blackspace will implement an automated [`cargo deny` check][urls.cargo_deny]. This
is a security checker for golang source code.

##### Fuzz Testing

Blackspace will implement automated fuzz testing to probe our code for other sources
of potential vulnerabilities.

## Building & Releasing

Blackspace takes care to secure the build and release process to prevent unintended
modifications.

### Network Security

All network traffic is secured via TLS and SSH. This includes checking out
Blackspace's code from the relevant [protected branch](#protected-branches),
Docker image retrieval, and publishment of Blackspace's release artifacts.

### Runtime Isolation

All builds run in an isolated sandbox that is destroyed after each use.

### Asset Audit Logging

Changes to Blackspace's assets will be logged through Digital Ocean's blob storage audit logging feature.

### Asset Signatures & Checksums

All assets are signed with checksums allowing users to verify asset authenticity
upon download. This verifies that assets have not been modified at rest.

## Vulnerability Reporting

We deeply appreciate any effort to discover and disclose security
vulnerabilities responsibly.

If you would like to report a vulnerability or have any security concerns with
Blackspace, please e-mail yoanyombapro@gmail.com.

For non-critical matters, we prefer users [open an issue][urls.new_security_report].
For us to best investigate your request, please include any of the
following when reporting:

* Proof of concept
* Any tools, including versions used
* Any relevant output

We take all disclosures very seriously and will do our best to rapidly respond
and verify the vulnerability before taking the necessary steps to fix it. After
our initial reply to your disclosure, which should be directly after receiving
it, we will periodically update you with the status of the fix.


[urls.cargo_deny]: https://github.com/securego/gosec
[urls.git]: https://git-scm.com/
[urls.new_security_report]: https://github.com/Lens-Platform/Platform/issues/new
[urls.blackspace_chat]: https://gitter.im/BlackspaceInc/community
[urls.blackspace_issues]: https://github.com/Lens-Platform/Platform/issues
[urls.blackspace_pull_requests]: https://github.com/Lens-Platform/Platform/pulls
[urls.blackspace_repo]: https://github.com/Lens-Platform/Platform
[urls.blackspace_roadmap]: https://github.com/Lens-Platform/Platform/milestones
