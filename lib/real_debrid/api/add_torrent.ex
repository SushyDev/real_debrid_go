defmodule RealDebrid.Api.AddTorrent do
  @moduledoc """
  Add a torrent file to Real-Debrid.
  """

  alias RealDebrid.Client

  @type response :: %{
          id: String.t(),
          uri: String.t()
        }

  @doc """
  Add a torrent file to the user's torrents.

  ## Parameters

    - `client` - The Real-Debrid client
    - `torrent_data` - Binary content of the torrent file
    - `opts` - Optional keyword list:
      - `:filename` - The filename (defaults to "upload.torrent")
      - `:host` - The host to use (optional)

  ## Returns

    - `{:ok, %{id: id, uri: uri}}` on success
    - `{:error, reason}` on failure
  """
  @spec add(Client.t(), binary(), keyword()) :: {:ok, response()} | {:error, term()}
  def add(%Client{} = client, torrent_data, opts \\ []) do
    filename = Keyword.get(opts, :filename, "upload.torrent")
    host = Keyword.get(opts, :host)

    multipart = [
      {:file, torrent_data, filename: filename, content_type: "application/x-bittorrent"}
    ]

    multipart =
      case host do
        nil -> multipart
        h -> multipart ++ [{"host", h}]
      end

    case Client.put_multipart(client, "/torrents/addTorrent", multipart) do
      {:ok, body} ->
        {:ok,
         %{
           id: body["id"],
           uri: body["uri"]
         }}

      {:error, reason} ->
        {:error, reason}
    end
  end
end
