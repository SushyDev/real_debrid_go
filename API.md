# Real-Debrid API Documentation

## Implementation Details

*   **Base URL:** `https://api.real-debrid.com/rest/1.0/`
*   **Method Grouping:** Methods are grouped by namespaces (e.g., "unrestrict", "user").
*   **HTTP Verbs:** `GET`, `POST`, `PUT`, `DELETE`.
    *   *Override:* If your client does not support all verbs, use the `X-HTTP-Verb` header.
*   **Response Format:**
    *   Success: HTTP 200 with a JSON object (unless specified otherwise).
    *   Errors: HTTP 4XX or 5XX, returning a JSON object with `error` (message) and `error_code` (optional integer).
*   **Encoding:** All strings must be UTF-8 encoded. Normalize to Unicode Normalization Form C (NFC) before encoding.
*   **Headers:** The API sends `ETag` headers and supports `If-None-Match`.
*   **Date Format:** Dates follow the Javascript `date.toJSON` format.
*   **Rate Limiting:** 250 requests per minute. Refused requests return HTTP 429 and count towards the limit. Bruteforcing leads to blocking.

---

## API Methods

### General

#### `GET /disable_access_token`
**Description:** Disable current access token.
*   **Authentication:** Required
*   **Return Value:** None (Returns HTTP 204)
*   **Possible Error Codes:**
    | HTTP Code | Reason |
    | :--- | :--- |
    | 401 | Bad token (expired, invalid) |

#### `GET /time`
**Description:** Get server time.
*   **Authentication:** None
*   **Return Value:** `Y-m-d H:i:s` (Raw data)

#### `GET /time/iso`
**Description:** Get server time in ISO format.
*   **Authentication:** None
*   **Return Value:** `Y-m-dTH:i:sO` (Raw data)

---

### User (`/user`)

#### `GET /user`
**Description:** Returns information on the current user.
*   **Authentication:** Required
*   **Return Value:**
```json
{
  "id": int,
  "username": "string",
  "email": "string",
  "points": int, // Fidelity points
  "locale": "string", // User language
  "avatar": "string", // URL
  "type": "string", // "premium" or "free"
  "premium": int, // seconds left as a Premium user
  "expiration": "string" // jsonDate
}
```
*   **Possible Error Codes:**
    | HTTP Code | Reason |
    | :--- | :--- |
    | 401 | Bad token (expired, invalid) |
    | 403 | Permission denied (account locked) |

---

### Unrestrict (`/unrestrict`)

#### `POST /unrestrict/check`
**Description:** Check if a file is downloadable on the concerned hoster.
*   **Authentication:** None
*   **Parameters:**
    | Name | Type | Description |
    | :--- | :--- | :--- |
    | `link` * | string | The original hoster link |
    | `password` | string | Password to unlock the file access hoster side |
*   **Return Value:**
```json
{
  "host": "string", // Host main domain
  "link": "string",
  "filename": "string",
  "filesize": int,
  "supported": int
}
```
*   **Possible Error Codes:**
    | HTTP Code | Reason |
    | :--- | :--- |
    | 503 | File unavailable |

#### `POST /unrestrict/link`
**Description:** Unrestrict a hoster link and get a new unrestricted link.
*   **Authentication:** Required
*   **Parameters:**
    | Name | Type | Description |
    | :--- | :--- | :--- |
    | `link` * | string | The original hoster link |
    | `password` | string | Password to unlock the file access hoster side |
    | `remote` | int | 0 or 1, use Remote traffic, dedicated servers and account sharing protections lifted |
*   **Return Value (Unique Link):**
```json
{
  "id": "string",
  "filename": "string",
  "mimeType": "string", // Mime Type guessed by extension
  "filesize": int, // Bytes, 0 if unknown
  "link": "string", // Original link
  "host": "string", // Host main domain
  "chunks": int, // Max Chunks allowed
  "crc": int, // Disable / enable CRC check
  "download": "string", // Generated link
  "streamable": int // Is the file streamable on website
}
```
*   **Possible Error Codes:**
    | HTTP Code | Reason |
    | :--- | :--- |
    | 401 | Bad token |
    | 403 | Permission denied |

#### `POST /unrestrict/folder`
**Description:** Unrestrict a hoster folder link and get individual links. Returns empty array if no links found.
*   **Authentication:** Required
*   **Parameters:**
    | Name | Type | Description |
    | :--- | :--- | :--- |
    | `link` * | string | The hoster folder link |
*   **Return Value:** Array of link objects (see `/unrestrict/link` schema).
*   **Possible Error Codes:** 401, 403

#### `PUT /unrestrict/containerFile`
**Description:** Decrypt a container file (RSDF, CCF, CCF3, DLC).
*   **Authentication:** Required
*   **Return Value:** Array of link objects.
*   **Possible Error Codes:** 400 (Bad Request), 401, 403, 503

#### `POST /unrestrict/containerLink`
**Description:** Decrypt a container file from a link.
*   **Authentication:** Required
*   **Parameters:**
    | Name | Type | Description |
    | :--- | :--- | :--- |
    | `link` * | string | HTTP Link of the container file |
*   **Return Value:** Array of link objects.
*   **Possible Error Codes:** 400, 401, 403, 503

---

### Traffic (`/traffic`)

#### `GET /traffic`
**Description:** Get traffic information for limited hosters (limits, current usage, extra packages).
*   **Authentication:** Required
*   **Return Value:**
```json
{
  "string": { // Host main domain
    "left": int, // Available bytes / links to use
    "bytes": int, // Bytes downloaded
    "links": int, // Links unrestricted
    "limit": int,
    "type": "string", // "links", "gigabytes", "bytes"
    "extra": int, // Additional traffic / links bought
    "reset": "string" // "daily", "weekly" or "monthly"
  }
}
```
*   **Possible Error Codes:** 401, 403

#### `GET /traffic/details`
**Description:** Get traffic details on each hoster used during a defined period.
*   **Authentication:** Required
*   **Parameters:**
    | Name | Type | Description |
    | :--- | :--- | :--- |
    | `start` | date | Start period (YYYY-MM-DD), default: a week ago |
    | `end` | date | End period (YYYY-MM-DD), default: today |
*   **Warning:** Period cannot exceed 31 days.
*   **Return Value:**
```json
{
  "YYYY-MM-DD": {
    "host": { // By Host main domain
      "string": int, // bytes downloaded on concerned host
    },
    "bytes": int // Total downloaded (in bytes) this day
  }
}
```
*   **Possible Error Codes:** 401, 403

---

### Streaming (`/streaming`)

#### `GET /streaming/transcode/{id}`
**Description:** Get transcoding links for given file. `{id}` from `/downloads` or `/unrestrict/link`.
*   **Authentication:** Required
*   **Return Value:**
```json
{
  "apple": { "quality": "string" }, // M3U8 Live Streaming
  "dash": { "quality": "string" }, // MPD Live Streaming
  "liveMP4": { "quality": "string" },
  "h264WebM": { "quality": "string" }
}
```
*   **Possible Error Codes:** 401, 403

#### `GET /streaming/mediaInfos/{id}`
**Description:** Get detailed media information for given file. `{id}` from `/downloads` or `/unrestrict/link`.
*   **Authentication:** Required
*   **Return Value:**
```json
{
  "filename": "string",
  "hoster": "string",
  "link": "string",
  "type": "string", // "movie", "show", "audio"
  "season": "string", // or null
  "episode": "string", // or null
  "year": "string", // or null
  "duration": float,
  "bitrate": int,
  "size": int,
  "details": {
    "video": { "und1": { "stream": "string", "lang": "string", "codec": "string", "width": int, "height": int, ... } },
    "audio": { "und1": { "stream": "string", "lang": "string", "codec": "string", "channels": float, ... } },
    "subtitles": [ "und1": { "lang": "string", "type": "string", ... } ]
  },
  "poster_path": "string",
  "audio_image": "string",
  "backdrop_path": "string"
}
```
*   **Possible Error Codes:** 401, 403, 503

---

### Downloads (`/downloads`)

#### `GET /downloads`
**Description:** Get user downloads list.
*   **Authentication:** Required
*   **Parameters:**
    | Name | Type | Description |
    | :--- | :--- | :--- |
    | `offset` | int | Starting offset |
    | `page` | int | Pagination system (prioritized over offset) |
    | `limit` | int | Entries per page (0-5000, default 100) |
*   **Return Value:** List of download objects.
*   **Possible Error Codes:** 401, 403

#### `DELETE /downloads/delete/{id}`
**Description:** Delete a link from downloads list.
*   **Authentication:** Required
*   **Return Value:** None (HTTP 204)
*   **Possible Error Codes:** 401, 403, 404

---

### Torrents (`/torrents`)

#### `GET /torrents`
**Description:** Get user torrents list.
*   **Authentication:** Required
*   **Parameters:**
    | Name | Type | Description |
    | :--- | :--- | :--- |
    | `offset` | int | Starting offset |
    | `page` | int | Pagination system |
    | `limit` | int | Entries per page (0-5000, default 100) |
    | `filter` | string | "active" to list active torrents only |
*   **Return Value:**
```json
[
  {
    "id": "string",
    "filename": "string",
    "hash": "string",
    "bytes": int,
    "host": "string",
    "split": int,
    "progress": int,
    "status": "downloaded", // magnet_error, waiting_files_selection, downloading, etc.
    "added": "string", // jsonDate
    "links": [ "string" ],
    "ended": "string", // Only if finished
    "speed": int, // Only if downloading/compressing/uploading
    "seeders": int // Only if downloading/magnet_conversion
  }
]
```
*   **Possible Error Codes:** 401, 403

#### `GET /torrents/info/{id}`
**Description:** Get all information on the asked torrent.
*   **Authentication:** Required
*   **Return Value:** Similar to `/torrents` item but includes `original_filename`, `original_bytes`, and `files` array (paths, bytes, selection status).
*   **Possible Error Codes:** 401, 403

#### `GET /torrents/activeCount`
**Description:** Get currently active torrents number and maximum limit.
*   **Authentication:** Required
*   **Return Value:** `{ "nb": int, "limit": int }`
*   **Possible Error Codes:** 401, 403

#### `GET /torrents/availableHosts`
**Description:** Get available hosts to upload the torrent to.
*   **Authentication:** Required
*   **Return Value:** `[ { "host": "string", "max_file_size": int } ]`
*   **Possible Error Codes:** 401, 403

#### `PUT /torrents/addTorrent`
**Description:** Add a torrent file to download.
*   **Authentication:** Required
*   **Parameters:** `host` (string) - Hoster domain.
*   **Return Value:** `{ "id": "string", "uri": "string" }` (HTTP 201)
*   **Possible Error Codes:** 400, 401, 403, 503

#### `POST /torrents/addMagnet`
**Description:** Add a magnet link to download.
*   **Authentication:** Required
*   **Parameters:**
    | Name | Type | Description |
    | :--- | :--- | :--- |
    | `magnet` * | string | Magnet link |
    | `host` | string | Hoster domain |
*   **Return Value:** `{ "id": "string", "uri": "string" }` (HTTP 201)
*   **Possible Error Codes:** 400, 401, 403, 503

#### `POST /torrents/selectFiles/{id}`
**Description:** Select files of a torrent to start it.
*   **Authentication:** Required
*   **Parameters:**
    | Name | Type | Description |
    | :--- | :--- | :--- |
    | `files` * | string | Selected files IDs (comma separated) or "all" |
*   **Return Value:** None (HTTP 204)
*   **Possible Error Codes:** 202, 400, 401, 403, 404

#### `DELETE /torrents/delete/{id}`
**Description:** Delete a torrent from torrents list.
*   **Authentication:** Required
*   **Return Value:** None (HTTP 204)
*   **Possible Error Codes:** 401, 403, 404

---

### Hosts (`/hosts`)

#### `GET /hosts`
**Description:** Get supported hosts.
*   **Authentication:** None
*   **Return Value:** `{ "domain": { "id": "string", "name": "string", "image": "string" } }`

#### `GET /hosts/status`
**Description:** Get status of hosters and their status on competitors.
*   **Authentication:** None
*   **Return Value:**
```json
{
  "domain": {
    "id": "string",
    "name": "string",
    "image": "string",
    "supported": int,
    "status": "string", // "up", "down", "unsupported"
    "check_time": "string",
    "competitors_status": { ... }
  }
}
```

#### `GET /hosts/regex`
**Description:** Get all supported links Regex.
*   **Authentication:** None
*   **Return Value:** List of Regex strings.

#### `GET /hosts/regexFolder`
**Description:** Get all supported folder Regex.
*   **Authentication:** None
*   **Return Value:** List of Regex strings.

#### `GET /hosts/domains`
**Description:** Get all supported domains.
*   **Authentication:** None
*   **Return Value:** List of domains.

---

### Settings (`/settings`)

#### `GET /settings`
**Description:** Get current user settings with possible values to update.
*   **Authentication:** Required
*   **Return Value:** JSON object with `download_ports`, `locales`, `streaming_qualities`, etc.
*   **Possible Error Codes:** 401, 403

#### `POST /settings/update`
**Description:** Update a user setting.
*   **Authentication:** Required
*   **Parameters:**
    | Name | Type | Description |
    | :--- | :--- | :--- |
    | `setting_name` * | string | `download_port`, `locale`, `streaming_quality`, etc. |
    | `setting_value` * | string | Value from `/settings` |
*   **Return Value:** None (HTTP 204)
*   **Possible Error Codes:** 400, 401, 403

#### `POST /settings/convertPoints`
**Description:** Convert fidelity points.
*   **Authentication:** Required
*   **Return Value:** None (HTTP 204)
*   **Possible Error Codes:** 401, 403, 503

#### `POST /settings/changePassword`
**Description:** Send verification email to change password.
*   **Authentication:** Required
*   **Return Value:** None (HTTP 204)
*   **Possible Error Codes:** 401, 403

#### `PUT /settings/avatarFile`
**Description:** Upload a new user avatar image.
*   **Authentication:** Required
*   **Return Value:** None (HTTP 204)
*   **Possible Error Codes:** 400, 401, 403

#### `DELETE /settings/avatarDelete`
**Description:** Reset user avatar to default.
*   **Authentication:** Required
*   **Return Value:** None (HTTP 204)
*   **Possible Error Codes:** 401, 403

---

## Authentication

Calls requiring authentication expect an HTTP header:
`Authorization: Bearer your_api_token`

Alternatively (but less secure), via URL parameter: `?auth_token=your_api_token`

**Base OAuth2 URL:** `https://api.real-debrid.com/oauth/v2/`

### Workflow for Websites (3-Legged OAuth2)
1.  **Authorize Endpoint:** `/auth`
    *   Redirect user to: `https://api.real-debrid.com/oauth/v2/auth?client_id=ID&redirect_uri=URI&response_type=code&state=STRING`
2.  **Callback:** User is redirected to `redirect_uri` with `code`.
3.  **Token Endpoint:** `/token`
    *   POST request: `client_id`, `client_secret`, `code`, `redirect_uri`, `grant_type="authorization_code"`.
    *   Returns: `access_token`, `expires_in`, `refresh_token`.

### Workflow for Mobile Apps (Device Flow)
1.  **Device Endpoint:** `/device/code`
    *   GET request with `client_id`.
    *   Returns: `device_code`, `user_code`, `verification_url`, `interval`, `expires_in`.
2.  **Verification:** User enters `user_code` at `verification_url`.
3.  **Polling:** Poll `/token` every `interval` seconds.
    *   POST parameters: `client_id`, `client_secret`, `code` (value of device_code), `grant_type="http://oauth.net/grant_type/device/1.0"`.
    *   Returns: `access_token` once authorized.

### Workflow for Opensource Apps
Similar to Mobile Apps but generates unique credentials per user.
1.  **Device Endpoint:** `/device/code`
    *   GET request: `client_id`, `new_credentials=yes`.
2.  **Credentials Endpoint:** Poll `/device/credentials`
    *   Returns new `client_id` and `client_secret` bound to the user.
3.  **Token Endpoint:** `/token`
    *   Use the new credentials to request token via Device Flow.

---

## Error Codes

| Code | Message |
| :--- | :--- |
| -1 | Internal error |
| 1 | Missing parameter |
| 2 | Bad parameter value |
| 3 | Unknown method |
| 4 | Method not allowed |
| 5 | Slow down |
| 6 | Ressource unreachable |
| 7 | Resource not found |
| 8 | Bad token |
| 9 | Permission denied |
| 10 | Two-Factor authentication needed |
| 11 | Two-Factor authentication pending |
| 12 | Invalid login |
| 13 | Invalid password |
| 14 | Account locked |
| 15 | Account not activated |
| 16 | Unsupported hoster |
| 17 | Hoster in maintenance |
| 18 | Hoster limit reached |
| 19 | Hoster temporarily unavailable |
| 20 | Hoster not available for free users |
| 21 | Too many active downloads |
| 22 | IP Address not allowed |
| 23 | Traffic exhausted |
| 24 | File unavailable |
| 25 | Service unavailable |
| 26 | Upload too big |
| 27 | Upload error |
| 28 | File not allowed |
| 29 | Torrent too big |
| 30 | Torrent file invalid |
| 31 | Action already done |
| 32 | Image resolution error |
| 33 | Torrent already active |
| 34 | Too many requests |
| 35 | Infringing file |
| 36 | Fair Usage Policy |

---

### Additional Notes

*   **Changelog:** The API is currently in version 1.0. Any major changes will be reflected in the version number in the base URL.
*   **Support:** For implementation questions, refer to the official Real-Debrid support forums or contact support directly if you encounter server-side bugs.
