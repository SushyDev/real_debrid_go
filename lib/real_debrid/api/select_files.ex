defmodule RealDebrid.Api.SelectFiles do
  @moduledoc """
  Select files from a torrent for download.
  """

  alias RealDebrid.Client

  @doc """
  Select files from a torrent to download.

  ## Parameters

    - `client` - The Real-Debrid client
    - `torrent_id` - The torrent ID
    - `file_ids` - Comma-separated list of file IDs to select (e.g., "1,2,3" or "all")

  ## Returns

    - `:ok` on success (204 response)
    - `{:error, reason}` on failure
  """
  @spec select(Client.t(), String.t(), String.t()) :: :ok | {:error, term()}
  def select(%Client{} = client, torrent_id, file_ids) do
    form = %{"files" => file_ids}

    case Client.post(client, "/torrents/selectFiles/#{torrent_id}", form: form, expected_status: 204) do
      {:ok, _body} ->
        :ok

      {:error, reason} ->
        {:error, reason}
    end
  end
end
