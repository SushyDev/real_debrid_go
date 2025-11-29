defmodule RealDebrid.Api.Streaming do
  @moduledoc """
  Streaming-related API functions.
  """

  alias RealDebrid.Client

  @type streaming_transcode :: %{
          apple: map(),
          dash: map(),
          live_mp4: map(),
          h264_webm: map()
        }

  @type video_detail :: %{
          stream: String.t(),
          lang: String.t(),
          lang_iso: String.t(),
          codec: String.t(),
          colorspace: String.t(),
          width: integer(),
          height: integer()
        }

  @type audio_detail :: %{
          stream: String.t(),
          lang: String.t(),
          lang_iso: String.t(),
          codec: String.t(),
          sampling: integer(),
          channels: float()
        }

  @type subtitle_detail :: %{
          stream: String.t(),
          lang: String.t(),
          lang_iso: String.t(),
          type: String.t()
        }

  @type media_infos :: %{
          filename: String.t(),
          hoster: String.t(),
          link: String.t(),
          type: String.t(),
          season: String.t() | nil,
          episode: String.t() | nil,
          year: String.t() | nil,
          duration: float(),
          bitrate: integer(),
          size: integer(),
          details: %{
            video: map(),
            audio: map(),
            subtitles: map()
          },
          poster_path: String.t(),
          audio_image: String.t(),
          backdrop_path: String.t()
        }

  @doc """
  Get transcoding options for a media file.

  ## Parameters

    - `client` - The Real-Debrid client
    - `id` - The file ID

  ## Returns

    - `{:ok, transcode}` on success
    - `{:error, reason}` on failure
  """
  @spec get_transcode(Client.t(), String.t()) :: {:ok, streaming_transcode()} | {:error, term()}
  def get_transcode(%Client{} = client, id) do
    case Client.get(client, "/streaming/transcode/#{id}") do
      {:ok, body, _headers} ->
        {:ok,
         %{
           apple: body["apple"] || %{},
           dash: body["dash"] || %{},
           live_mp4: body["liveMP4"] || %{},
           h264_webm: body["h264WebM"] || %{}
         }}

      {:error, reason} ->
        {:error, reason}
    end
  end

  @doc """
  Get media information for a file.

  ## Parameters

    - `client` - The Real-Debrid client
    - `id` - The file ID

  ## Returns

    - `{:ok, media_infos}` on success
    - `{:error, reason}` on failure
  """
  @spec get_media_infos(Client.t(), String.t()) :: {:ok, media_infos()} | {:error, term()}
  def get_media_infos(%Client{} = client, id) do
    case Client.get(client, "/streaming/mediaInfos/#{id}") do
      {:ok, body, _headers} ->
        {:ok,
         %{
           filename: body["filename"],
           hoster: body["hoster"],
           link: body["link"],
           type: body["type"],
           season: body["season"],
           episode: body["episode"],
           year: body["year"],
           duration: body["duration"],
           bitrate: body["bitrate"],
           size: body["size"],
           details: %{
             video: body["details"]["video"] || %{},
             audio: body["details"]["audio"] || %{},
             subtitles: body["details"]["subtitles"] || %{}
           },
           poster_path: body["poster_path"],
           audio_image: body["audio_image"],
           backdrop_path: body["backdrop_path"]
         }}

      {:error, reason} ->
        {:error, reason}
    end
  end
end
