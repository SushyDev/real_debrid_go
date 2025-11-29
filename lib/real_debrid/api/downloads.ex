defmodule RealDebrid.Api.Downloads do
  @moduledoc """
  Downloads management API functions.
  """

  alias RealDebrid.Client

  @type download :: %{
          id: String.t(),
          filename: String.t(),
          mime_type: String.t(),
          filesize: integer(),
          link: String.t(),
          host: String.t(),
          chunks: integer(),
          crc: integer(),
          download: String.t(),
          generated: String.t(),
          type: String.t() | nil
        }

  @doc """
  Get downloads list with total count.

  ## Parameters

    - `client` - The Real-Debrid client
    - `opts` - Optional keyword list:
      - `:limit` - Number of results per page
      - `:page` - Page number
      - `:offset` - Offset for results

  ## Returns

    - `{:ok, %{downloads: downloads, total_count: count}}` on success
    - `{:error, reason}` on failure
  """
  @spec get(Client.t(), keyword()) :: {:ok, %{downloads: [download()], total_count: integer()}} | {:error, term()}
  def get(%Client{} = client, opts \\ []) do
    params =
      []
      |> maybe_add_param(:limit, Keyword.get(opts, :limit))
      |> maybe_add_param(:page, Keyword.get(opts, :page))
      |> maybe_add_param(:offset, Keyword.get(opts, :offset))
      |> Map.new()

    case Client.get(client, "/downloads", params: params) do
      {:ok, body, headers} ->
        downloads = Enum.map(body, &parse_download/1)

        total_count =
          headers
          |> get_header("x-total-count")
          |> parse_integer(0)

        {:ok, %{downloads: downloads, total_count: total_count}}

      {:error, reason} ->
        {:error, reason}
    end
  end

  @doc """
  Delete a download by its ID.

  ## Returns

    - `:ok` on success
    - `{:error, reason}` on failure
  """
  @spec delete(Client.t(), String.t()) :: :ok | {:error, term()}
  def delete(%Client{} = client, id) do
    Client.delete(client, "/downloads/delete/#{id}")
  end

  defp parse_download(data) do
    %{
      id: data["id"],
      filename: data["filename"],
      mime_type: data["mimeType"],
      filesize: data["filesize"],
      link: data["link"],
      host: data["host"],
      chunks: data["chunks"],
      crc: data["crc"],
      download: data["download"],
      generated: data["generated"],
      type: data["type"]
    }
  end

  defp maybe_add_param(params, _key, nil), do: params
  defp maybe_add_param(params, key, value), do: [{key, value} | params]

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
