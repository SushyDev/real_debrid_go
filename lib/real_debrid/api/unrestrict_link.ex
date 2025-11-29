defmodule RealDebrid.Api.UnrestrictLink do
  @moduledoc """
  Unrestrict a single link.
  """

  alias RealDebrid.Client

  @type response :: %{
          id: String.t(),
          filename: String.t(),
          mime_type: String.t(),
          filesize: integer(),
          link: String.t(),
          host: String.t(),
          chunks: integer(),
          crc: integer(),
          download: String.t(),
          streamable: integer()
        }

  @doc """
  Unrestrict a single link.

  ## Parameters

    - `client` - The Real-Debrid client
    - `link` - The link to unrestrict
    - `opts` - Optional keyword list:
      - `:password` - Password for the link
      - `:remote` - 0 or 1 for remote traffic

  ## Returns

    - `{:ok, response}` on success
    - `{:error, reason}` on failure
  """
  @spec unrestrict(Client.t(), String.t(), keyword()) :: {:ok, response()} | {:error, term()}
  def unrestrict(%Client{} = client, link, opts \\ []) do
    form = %{"link" => link}

    form =
      case Keyword.get(opts, :password) do
        nil -> form
        pwd -> Map.put(form, "password", pwd)
      end

    form =
      case Keyword.get(opts, :remote) do
        nil -> form
        remote -> Map.put(form, "remote", to_string(remote))
      end

    case Client.post(client, "/unrestrict/link", form: form) do
      {:ok, body} when is_map(body) ->
        {:ok,
         %{
           id: body["id"],
           filename: body["filename"],
           mime_type: body["mimeType"],
           filesize: body["filesize"],
           link: body["link"],
           host: body["host"],
           chunks: body["chunks"],
           crc: body["crc"],
           download: body["download"],
           streamable: body["streamable"]
         }}

      {:ok, body} when is_list(body) ->
        # Handle array response (folder items)
        items =
          Enum.map(body, fn item ->
            %{
              id: item["id"],
              filename: item["filename"],
              mime_type: item["mimeType"],
              filesize: item["filesize"],
              link: item["link"],
              host: item["host"],
              chunks: item["chunks"],
              download: item["download"],
              generated: item["generated"],
              type: item["type"]
            }
          end)

        {:ok, items}

      {:error, reason} ->
        {:error, reason}
    end
  end
end
