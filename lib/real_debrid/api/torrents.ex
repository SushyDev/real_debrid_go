defmodule RealDebrid.Api.Torrents do
  @moduledoc """
  Torrents listing API functions.
  """

  alias RealDebrid.Client

  @type torrent :: %{
          id: String.t(),
          filename: String.t(),
          hash: String.t(),
          bytes: integer(),
          host: String.t(),
          split: integer(),
          progress: integer(),
          status: String.t(),
          added: String.t(),
          links: [String.t()],
          ended: String.t() | nil,
          speed: integer() | nil,
          seeders: integer() | nil
        }

  @type torrents_response :: %{
          torrents: [torrent()],
          total_count: integer(),
          current_page: integer()
        }

  @doc """
  Get list of user's torrents.

  ## Parameters

    - `client` - The Real-Debrid client
    - `opts` - Optional keyword list:
      - `:limit` - Number of results per page (1-1000, default: 100)
      - `:page` - Page number (default: 1)

  ## Returns

    - `{:ok, %{torrents: torrents, total_count: count, current_page: page}}` on success
    - `{:error, reason}` on failure
  """
  @spec get(Client.t(), keyword()) :: {:ok, torrents_response()} | {:error, term()}
  def get(%Client{} = client, opts \\ []) do
    limit = Keyword.get(opts, :limit, 100)
    page = Keyword.get(opts, :page, 1)

    if limit < 1 or limit > 1000 do
      {:error, "limit must be between 1 and 1000, got #{limit}"}
    else
      params = %{limit: limit, page: page}

      case Client.get(client, "/torrents", params: params) do
        {:ok, body, headers} ->
          torrents = Enum.map(body, &parse_torrent/1)

          total_count =
            headers
            |> get_header("x-total-count")
            |> parse_integer(0)

          {:ok,
           %{
             torrents: torrents,
             total_count: total_count,
             current_page: page
           }}

        {:error, reason} ->
          {:error, reason}
      end
    end
  end

  @doc """
  Get all torrents (handles pagination automatically).

  ## Parameters

    - `client` - The Real-Debrid client

  ## Returns

    - `{:ok, torrents}` - List of all torrents
    - `{:error, reason}` on failure
  """
  @spec get_all(Client.t()) :: {:ok, [torrent()]} | {:error, term()}
  def get_all(%Client{} = client) do
    case get(client, limit: 1000, page: 1) do
      {:ok, %{torrents: torrents, total_count: total_count}} ->
        total_pages = div(total_count + 999, 1000)
        fetch_all_pages(client, torrents, 2, total_pages)

      {:error, reason} ->
        {:error, reason}
    end
  end

  @doc """
  Find a torrent by hash.

  ## Parameters

    - `torrents` - List of torrents
    - `hash` - The torrent hash to find

  ## Returns

    - `torrent` if found
    - `nil` if not found
  """
  @spec get_by_hash([torrent()], String.t()) :: torrent() | nil
  def get_by_hash(torrents, hash) do
    Enum.find(torrents, fn torrent -> torrent.hash == hash end)
  end

  defp fetch_all_pages(_client, torrents, page, total_pages) when page > total_pages do
    {:ok, torrents}
  end

  defp fetch_all_pages(client, torrents, page, total_pages) do
    case get(client, limit: 1000, page: page) do
      {:ok, %{torrents: new_torrents}} ->
        fetch_all_pages(client, torrents ++ new_torrents, page + 1, total_pages)

      {:error, reason} ->
        {:error, reason}
    end
  end

  defp parse_torrent(body) do
    %{
      id: body["id"],
      filename: body["filename"],
      hash: body["hash"],
      bytes: body["bytes"],
      host: body["host"],
      split: body["split"],
      progress: body["progress"],
      status: body["status"],
      added: body["added"],
      links: body["links"] || [],
      ended: body["ended"],
      speed: body["speed"],
      seeders: body["seeders"]
    }
  end

  defp get_header(headers, key) do
    case List.keyfind(headers, key, 0) do
      {_, value} -> value
      nil -> nil
    end
  end

  defp parse_integer(nil, default), do: default

  defp parse_integer(value, default) when is_binary(value) do
    case Integer.parse(value) do
      {int, _} -> int
      :error -> default
    end
  end

  defp parse_integer(value, _default) when is_integer(value), do: value
end
