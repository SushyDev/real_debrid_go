defmodule RealDebrid.Api.TorrentInfo do
  @moduledoc """
  Get detailed torrent information.
  """

  alias RealDebrid.Client

  @type torrent_file :: %{
          id: integer(),
          path: String.t(),
          bytes: integer(),
          selected: integer()
        }

  @type torrent_info :: %{
          id: String.t(),
          filename: String.t(),
          original_filename: String.t(),
          hash: String.t(),
          bytes: integer(),
          original_bytes: integer(),
          host: String.t(),
          split: integer(),
          progress: integer(),
          status: String.t(),
          added: String.t(),
          files: [torrent_file()],
          links: [String.t()],
          ended: String.t() | nil,
          speed: integer() | nil,
          seeders: integer() | nil
        }

  @doc """
  Get detailed information about a torrent.

  ## Parameters

    - `client` - The Real-Debrid client
    - `id` - The torrent ID

  ## Returns

    - `{:ok, torrent_info}` on success
    - `{:error, reason}` on failure
  """
  @spec get(Client.t(), String.t()) :: {:ok, torrent_info()} | {:error, term()}
  def get(%Client{} = client, id) do
    case Client.get(client, "/torrents/info/#{id}") do
      {:ok, body, _headers} ->
        {:ok, parse_torrent_info(body)}

      {:error, reason} ->
        {:error, reason}
    end
  end

  defp parse_torrent_info(body) do
    %{
      id: body["id"],
      filename: body["filename"],
      original_filename: body["original_filename"],
      hash: body["hash"],
      bytes: body["bytes"],
      original_bytes: body["original_bytes"],
      host: body["host"],
      split: body["split"],
      progress: body["progress"],
      status: body["status"],
      added: body["added"],
      files: parse_files(body["files"] || []),
      links: body["links"] || [],
      ended: body["ended"],
      speed: body["speed"],
      seeders: body["seeders"]
    }
  end

  defp parse_files(files) do
    Enum.map(files, fn file ->
      %{
        id: file["id"],
        path: file["path"],
        bytes: file["bytes"],
        selected: file["selected"]
      }
    end)
  end
end
