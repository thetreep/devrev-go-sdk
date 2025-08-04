## v0.2.2 (2025-08-04)

### Public API Changes

#### Changes from 2025-07-29

## New Endpoints

### Vistas
  - Initial support for `meetings` vistas

## Modified Endpoints

### Articles
- Added field `parent` to [`/articles.list`](/public/api-reference/articles/list-articles-post)

### Customization
- Added fields to [`/schemas.aggregated.get`](/public/api-reference/customization/aggregated-schema-get):
  - `custom_schema_fragments`
  - `stock_schema_fragment`

<!-- included from 2025-07-29 -->

### Beta API Changes

#### Changes from 2025-07-29

## New Endpoints

### Vistas
- [`/vistas.create`](/beta/api-reference/vistas/create): Creates a vista
  - Initial support for `meetings` vistas

## Modified Endpoints

### Articles
- Added field `parent` to [`/articles.list`](/beta/api-reference/articles/list-articles-post)
- Added field `parent` to [`/articles.count`](/beta/api-reference/articles/count-post)

### Customization
- Added fields to [`/schemas.aggregated.get`](/beta/api-reference/customization/aggregated-schema-get):
  - `custom_schema_fragments`
  - `stock_schema_fragment`

<!-- included from 2025-07-29 -->

## v0.2.1 (2025-06-30)

### Public API Changes

#### Changes from 2025-06-24

## Promoted Endpoints

### Custom Objects
- [`/custom-objects.count`](/public/api-reference/customization/custom-objects-count)
- [`/custom-objects.create`](/public/api-reference/customization/custom-objects-create)
- [`/custom-objects.delete`](/public/api-reference/customization/custom-objects-delete)
- [`/custom-objects.get`](/public/api-reference/customization/custom-objects-get-post)
- [`/custom-objects.list`](/public/api-reference/customization/custom-objects-list-post)
- [`/custom-objects.update`](/public/api-reference/customization/custom-objects-update)

### Custom Stages
- [`/stages.custom.create`](/public/api-reference/customization/custom-stages-create)
- [`/stages.custom.get`](/public/api-reference/customization/custom-stages-get-post)
- [`/stages.custom.list`](/public/api-reference/customization/custom-stages-list-post)
- [`/stages.custom.update`](/public/api-reference/customization/custom-stages-update)

### Custom States
- [`/states.custom.create`](/public/api-reference/customization/custom-states-create)
- [`/states.custom.get`](/public/api-reference/customization/custom-states-get-post)
- [`/states.custom.list`](/public/api-reference/customization/custom-states-list-post)
- [`/states.custom.update`](/public/api-reference/customization/custom-states-update)

### Customization
- [`/schemas.aggregated.get`](/public/api-reference/customization/aggregated-schema-get)
- [`/schemas.custom.get`](/public/api-reference/customization/custom-schema-fragments-get-post)
- [`/schemas.custom.list`](/public/api-reference/customization/custom-schema-fragments-list-post)
- [`/schemas.custom.set`](/public/api-reference/customization/custom-schema-fragments-set)
- [`/schemas.stock.get`](/public/api-reference/customization/stock-schema-fragments-get-post)
- [`/schemas.stock.list`](/public/api-reference/customization/stock-schema-fragments-list-post)
- [`/schemas.subtypes.prepare-update`](/public/api-reference/customization/schemas-subtype-prepare-update-get)

### Stage Diagrams
- [`/stage-diagrams.create`](/public/api-reference/customization/stage-diagrams-create)
- [`/stage-diagrams.get`](/public/api-reference/customization/stage-diagrams-get-post)
- [`/stage-diagrams.list`](/public/api-reference/customization/stage-diagrams-list-post)
- [`/stage-diagrams.update`](/public/api-reference/customization/stage-diagrams-update)

<!-- included from 2025-06-24 -->

### Beta API Changes

#### Changes from 2025-06-24

## New Endpoints

### Vistas
- [`/vistas.delete`](/beta/api-reference/vistas/delete): Deletes a vista
- [`/vistas.get`](/beta/api-reference/vistas/get-post): Gets a vista's information
- [`/vistas.list`](/beta/api-reference/vistas/list-post): Lists available vistas

### Vista Group Items
- [`/vistas.groups.delete`](/beta/api-reference/vistas/groups-delete): Deletes a vista group item
- [`/vistas.groups.get`](/beta/api-reference/vistas/groups-get-post): Gets a vista group item's information
- [`/vistas.groups.list`](/beta/api-reference/vistas/groups-list-post): Lists the available vista group items

## Updated Endpoints

### Webhooks
- Added webhook types for articles:
  - `article_created`
  - `article_updated`
  - `article_deleted`

<!-- included from 2025-06-24 -->

## v0.2.0 (2025-06-05)

### Public API Changes

### Beta API Changes

## v0.1.1 (2025-06-04)

### Public API Changes

#### Changes from 2025-06-02

## Promoted Endpoints

### Accounts
- [`/accounts.merge`](/public/api-reference/accounts/merge)

### Chats
- [`/chats.create`](/public/api-reference/chats/create)
- [`/chats.get`](/public/api-reference/chats/get-post)

### Code Changes
- [`/code-changes.create`](/public/api-reference/code-changes/create)
- [`/code-changes.delete`](/public/api-reference/code-changes/delete)
- [`/code-changes.get`](/public/api-reference/code-changes/get-post)
- [`/code-changes.list`](/public/api-reference/code-changes/list-post)
- [`/code-changes.update`](/public/api-reference/code-changes/update)

### Commands
- [`/commands.create`](/public/api-reference/commands/create)
- [`/commands.get`](/public/api-reference/commands/get-post)
- [`/commands.list`](/public/api-reference/commands/list-post)
- [`/commands.update`](/public/api-reference/commands/update)

### Dev Organizations
- [`/dev-orgs.get`](/public/api-reference/dev-orgs/get-post)

### Directories
- [`/directories.count`](/public/api-reference/directory/directories-count)
- [`/directories.create`](/public/api-reference/directory/directories-create)
- [`/directories.delete`](/public/api-reference/directory/directories-delete)
- [`/directories.get`](/public/api-reference/directory/directories-get-post)
- [`/directories.list`](/public/api-reference/directory/directories-list-post)
- [`/directories.update`](/public/api-reference/directory/directories-update)

### Meetings
- [`/meetings.count`](/public/api-reference/meetings/count)
- [`/meetings.create`](/public/api-reference/meetings/create)
- [`/meetings.delete`](/public/api-reference/meetings/delete)
- [`/meetings.get`](/public/api-reference/meetings/get-post)
- [`/meetings.list`](/public/api-reference/meetings/list-post)
- [`/meetings.update`](/public/api-reference/meetings/update)

### Metric Definitions
- [`/metric-definitions.create`](/public/api-reference/slas/metric-definitions-create)
- [`/metric-definitions.delete`](/public/api-reference/slas/metric-definitions-delete)
- [`/metric-definitions.execute`](/public/api-reference/slas/metric-definitions-execute)
- [`/metric-definitions.get`](/public/api-reference/slas/metric-definitions-get-post)
- [`/metric-definitions.update`](/public/api-reference/slas/metric-definitions-update)

### Organization Schedules
- [`/org-schedules.create`](/public/api-reference/schedules/org-schedules-create)
- [`/org-schedules.get`](/public/api-reference/schedules/org-schedules-get-post)
- [`/org-schedules.list`](/public/api-reference/schedules/org-schedules-list-post)
- [`/org-schedules.set-future`](/public/api-reference/schedules/org-schedules-set-future)
- [`/org-schedules.transition`](/public/api-reference/schedules/org-schedules-transition)
- [`/org-schedules.update`](/public/api-reference/schedules/org-schedules-update)

### Organization Schedule Fragments
- [`/org-schedule-fragments.create`](/public/api-reference/schedules/org-schedule-fragments-create)
- [`/org-schedule-fragments.evaluate`](/public/api-reference/schedules/org-schedule-fragments-evaluate)
- [`/org-schedule-fragments.get`](/public/api-reference/schedules/org-schedule-fragments-get-post)
- [`/org-schedule-fragments.transition`](/public/api-reference/schedules/org-schedule-fragments-transition)

### Reactions
- [`/reactions.list`](/public/api-reference/timeline-entries/reactions-list-post)
- [`/reactions.update`](/public/api-reference/timeline-entries/reactions-update)

### Snap Widgets
- [`/snap-widgets.create`](/public/api-reference/snap-widgets/create)

### Web Crawler Jobs
- [`/web-crawler-jobs.control`](/public/api-reference/web-crawler-job/web-crawler-jobs-control)
- [`/web-crawler-jobs.create`](/public/api-reference/web-crawler-job/create-web-crawler-job)
- [`/web-crawler-jobs.get`](/public/api-reference/web-crawler-job/get-web-crawler-job-post)
- [`/web-crawler-jobs.list`](/public/api-reference/web-crawler-job/list-web-crawler-jobs-post)

## New Endpoints

### Chats
- [`/chats.update`](/public/api-reference/chats/update): Updates a chat (DM)

### Surveys
- [`/surveys.get`](/public/api-reference/surveys/get-post): Gets the information for a survey

## Modified Endpoints

### Accounts
- Adds `primary_account` field, which is set when the account is merged with another

### Rev Orgs
- Adds `primary_rev_org` field, which is set when the Rev org (workspace) is merged with another

### Users
- Adds `primary_identity` field, which is set when the user is merged with another

<!-- included from 2025-06-02 -->

#### Changes from 2025-05-19

## New Endpoints

### Artifacts
- [`/artifacts.versions.delete`](/public/api-reference/artifacts/hard-delete-version) - Delete artifact versions
- [`/artifacts.versions.prepare`](/public/api-reference/artifacts/versions-prepare) - Prepare artifact versions

### Links
- [`/links.replace`](/public/api-reference/links/replace) - Replace links

### Metric Trackers
- [`/metric-trackers.get`](/public/api-reference/slas/metric-trackers-get-post) - Get metric tracker details

## Modified Endpoints

### Accounts
Added `tier` filter support to:
- [`/accounts.export`](/public/api-reference/accounts/export-post)
- [`/accounts.list`](/public/api-reference/accounts/list-post)

### Articles
Modified schema for:
- [`/articles.create`](/public/api-reference/articles/create-article)
  - Made `applies_to_parts` optional (removed min items constraint)
  - Added `content_format` and `data_sources` properties
- [`/articles.update`](/public/api-reference/articles/update-article)
  - Added `content_format` and `data_sources` properties

### Conversations
- [`/conversations.list`](/public/api-reference/conversations/list-post) - Added `state` and `actual_close_date` filters
- [`/conversations.update`](/public/api-reference/conversations/update) - Made empty owner sets allowed (removed min items constraint)

### Groups
Added new filters to [`/groups.list`](/public/api-reference/groups/list-post):
- `created_by`
- `name`
- `sync_metadata` related filters

### SLAs
Added `tier` support to account selector in:
- [`/slas.create`](/public/api-reference/slas/create)
- [`/slas.update`](/public/api-reference/slas/update)

### Webhooks
Added new event type `sla_tracker_fetched` across webhook endpoints:
- [`/webhooks.create`](/public/api-reference/webhooks/create)
- [`/webhooks.event`](/public/api-reference/webhooks/event)
- [`/webhooks.get`](/public/api-reference/webhooks/get-post)
- [`/webhooks.list`](/public/api-reference/webhooks/list-post)
- [`/webhooks.update`](/public/api-reference/webhooks/update)

### Works
Added sync history filters to:
- [`/works.export`](/public/api-reference/works/export-post)
- [`/works.list`](/public/api-reference/works/list-post)

<!-- included from 2025-05-19 -->

#### Changes from 2025-04-24


## New Endpoints

### Service Accounts
- Added [`POST /service-accounts.create`](/public/api-reference/service-accounts/create) for creating service accounts

## Modified Endpoints

### Customization
- Added `is_overridable` property to schema enum fields

### Timeline
- Added `references` property to timeline comments

<!-- included from 2025-04-24 -->

#### Changes from 2025-04-13


# API Changelog

## Modified Endpoints

### SLAs
Added field `in_policy` to SLA metric target to indicate whether the metric is in the active SLA policy

### Webhooks
New event types added to webhook related endpoints:
- `dashboard_created`
- `dashboard_deleted`
- `dashboard_updated`
- `widget_created`
- `widget_deleted`
- `widget_updated`

<!-- included from 2025-04-13 -->

#### Changes from 2025-04-02


### Groups
- [`/groups.members.list`](/public/api-reference/groups/group-members-list-post): Added new property `member_rev_org` to response schema for both GET and POST methods

### Parts
Added new property `stage_validation_options` to request body for:
- [`/parts.create`](/public/api-reference/parts/create)
- [`/parts.update`](/public/api-reference/parts/update)

<!-- included from 2025-04-02 -->

### Beta API Changes

#### Changes from 2025-06-02

## New Endpoints

### Chats
- [`/chats.update`](/beta/api-reference/chats/update): Updates a chat (DM)

### Record Templates
- [`/record-templates.get`](/beta/api-reference/record-templates/record-template-get-post): Gets a record template's information

### Service Accounts
- [`/service-accounts.update`](/beta/api-reference/service-accounts/update): Updates a service account's information

### Subscribers
- [`/subscribers.list`](/beta/api-reference/subscribers/list-post): Lists subscribers for an object

### Surveys
- [`/surveys.get`](/beta/api-reference/surveys/get-post): Gets the information for a survey

## Modified Endpoints

### Accounts
- Adds `primary_account` field, which is set when the account is merged with another

### AI Agents
- Adds `suggestions` agent response type to the `ai_agent_response` webhook event

### Rev Orgs
- Adds `primary_rev_org` field, which is set when the Rev org (workspace) is merged with another

### Users
- Adds `primary_identity` field, which is set when the user is merged with another

<!-- included from 2025-06-02 -->

#### Changes from 2025-05-19

## New Endpoints

### Artifacts
- [`/artifacts.versions.delete`](/beta/api-reference/artifacts/hard-delete-version): Delete an artifact version

### Brands
- [`/brands.create`](/beta/api-reference/brands/create): Create a new brand
- [`/brands.delete`](/beta/api-reference/brands/delete): Delete an existing brand
- [`/brands.get`](/beta/api-reference/brands/get-post): Get the information for a brand
- [`/brands.list`](/beta/api-reference/brands/list-post): List brands matching a given filter
- [`/brands.update`](/beta/api-reference/brands/update): Update an existing brand

### Links
- [`/links.replace`](/beta/api-reference/links/replace): Replace a link

### Metrics
- [`/metric-trackers.get`](/beta/api-reference/slas/metric-trackers-get-post): Get a metric tracker

## Modified Endpoints

### Accounts
- Added `subtype` and `tier` query parameters to [`/accounts.export`](/beta/api-reference/accounts/export-post) and [`/accounts.list`](/beta/api-reference/accounts/list-post)
- Added `drop` property to `custom_schema_spec` in [`/accounts.create`](/beta/api-reference/accounts/create) and [`/accounts.update`](/beta/api-reference/accounts/update)

### Articles
- Added `content_format` and `data_sources` properties in [`/articles.create`](/beta/api-reference/articles/create-article) and [`/articles.update`](/beta/api-reference/articles/update-article)
- Made `applies_to_parts` optional in [`/articles.create`](/beta/api-reference/articles/create-article)

### Commands
- Added `ai_assistant_chat` and `dm` to allowed object types in surface configuration for [`/commands.create`](/beta/api-reference/commands/create), [`/commands.update`](/beta/api-reference/commands/update) and related endpoints

### Conversations
- Added `state` query parameter to [`/conversations.list`](/beta/api-reference/conversations/list-post) and [`/conversations.export`](/beta/api-reference/conversations/export-post)
- Modified `owned_by.set` to allow empty arrays in [`/conversations.update`](/beta/api-reference/conversations/update)

### Groups
- Added filtering by `created_by`, `name`, and `sync_metadata` fields to [`/groups.list`](/beta/api-reference/groups/list-post)

### Links
- Added `custom_link_type` property to link objects in [`/links.create`](/beta/api-reference/links/create) and related endpoints

### Meetings
- Added `sync_metadata` property to meeting objects in [`/meetings.create`](/beta/api-reference/meetings/create) and related endpoints

### Opportunities
- Added `contacts` property to opportunity objects in work-related endpoints

### SLAs
- Added `subtype` and `tier` properties to account selector in [`/slas.create`](/beta/api-reference/slas/create) and [`/slas.update`](/beta/api-reference/slas/update)

### Webhooks
- Added `sla_tracker_fetched` event type to [`/webhooks.create`](/beta/api-reference/webhooks/create), [`/webhooks.update`](/beta/api-reference/webhooks/update) and related endpoints

### Widgets
- Added `combination` visualization type
- Removed `percentage_precision` from PVP metric configuration

## Schema Changes

### Custom Schema Spec
Added `drop` property to `custom_schema_spec` in multiple endpoints including:
- Account endpoints
- Article endpoints
- Chat endpoints
- Code change endpoints
- Conversation endpoints
- Custom object endpoints
- Dev user endpoints
- Incident endpoints
- Meeting endpoints
- Part endpoints
- Work endpoints

### Meeting Summary
Added new subschema to meeting summary objects affecting multiple endpoints that reference meetings

### Link Summary
Modified link summary schema to support custom link types across multiple endpoints

<!-- included from 2025-05-19 -->

#### Changes from 2025-04-24


## New Endpoints

### Service Accounts
- Added [`POST /service-accounts.create`](/beta/api-reference/service-accounts/create) for creating service accounts

### Widgets
- Added [`GET /widgets.get`](/beta/api-reference/widgets/get) for retrieving widgets

## Modified Features

### Customization
- Added `is_overridable` property to schema enum fields

### Timeline
- Added `references` property to timeline comments

### Parts
- Added support for part types `linkable` and `runnable`

<!-- included from 2025-04-24 -->

#### Changes from 2025-04-13


# API Changelog

## Modified Endpoints

### SLAs
Added field `in_policy` to SLA metric target to indicate whether the metric is in the active SLA policy

### Webhooks
New event types added to webhook related endpoints:
- `dashboard_created`
- `dashboard_deleted`
- `dashboard_updated`
- `widget_created`
- `widget_deleted`
- `widget_updated`

<!-- included from 2025-04-13 -->

#### Changes from 2025-04-02


### Dev Users
- Added new endpoint [`/dev-users.merge`](/beta/api-reference/dev-users/merge) for merging Dev users

### Groups
- Added `member_rev_org` property to member objects returned by:
  - [`/groups.members.list`](/beta/api-reference/groups/group-members-list-post) (both GET and POST methods)

### Parts
- Added `stage_validation_options` property to request body of:
  - [`/parts.create`](/beta/api-reference/parts/create)
  - [`/parts.update`](/beta/api-reference/parts/update)
- Added `custom_fields` parameter to:
  - [`/parts.list`](/beta/api-reference/parts/list-post) (both GET and POST methods)

### Question Answers
- Added `sources` property to question-answer object

<!-- included from 2025-04-02 -->

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## v0.1.0 (2025-01-01)

### Added

- Initial release of DevRev Go SDK
- Support for Public API
- Support for Beta API
- Automated build and release workflow
