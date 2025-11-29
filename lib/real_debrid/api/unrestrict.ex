defmodule RealDebrid.Api.Unrestrict do
  @moduledoc """
  Unrestrict links API functions.
  """

  alias RealDebrid.Client

  @type check_response :: %{
          host: String.t(),
          link: String.t(),
          filename: String.t(),
          filesize: integer(),
          supported: integer()
        }

  @type folder_item :: %{
          id: String.t(),
          filename: String.t(),
          mime_type: String.t(),
          filesize: integer(),
          link: String.t(),
          host: String.t(),
          chunks: integer(),
          download: String.t(),
          generated: String.t(),
          type: String.t() | nil
        }

  @type link_options :: %{
          password: String.t() | nil,
          remote: integer() | nil
        }

  @doc """
  Check if a link is supported.

  ## Parameters

    - `client` - The Real-Debrid client
    - `link` - The link to check
    - `password` - Optional password for the link

  ## Returns

    - `{:ok, check_response}` on success
    - `{:error, reason}` on failure
  """
  @spec check(Client.t(), String.t(), String.t() | nil) ::
          {:ok, check_response()} | {:error, term()}
  def check(%Client{} = client, link, password \\ nil) do
    form =
      case password do
        nil -> %{"link" => link}
        pwd -> %{"link" => link, "password" => pwd}
      end

    case Client.post(client, "/unrestrict/check", form: form) do
      {:ok, body} ->
        {:ok,
         %{
           host: body["host"],
           link: body["link"],
           filename: body["filename"],
           filesize: body["filesize"],
           supported: body["supported"]
         }}

      {:error, reason} ->
        {:error, reason}
    end
  end

  @doc """
  Unrestrict a folder link.

  ## Parameters

    - `client` - The Real-Debrid client
    - `link` - The folder link to unrestrict

  ## Returns

    - `{:ok, items}` - List of folder items
    - `{:error, reason}` on failure
  """
  @spec folder(Client.t(), String.t()) :: {:ok, [folder_item()]} | {:error, term()}
  def folder(%Client{} = client, link) do
    form = %{"link" => link}

    case Client.post(client, "/unrestrict/folder", form: form) do
      {:ok, body} ->
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

  @doc """
  Unrestrict a container file (RSDF/CCF/DLC) by uploading it.

  ## Parameters

    - `client` - The Real-Debrid client
    - `file_data` - Binary content of the container file
    - `filename` - Name of the file

  ## Returns

    - `{:ok, links}` - List of unrestricted links
    - `{:error, reason}` on failure
  """
  @spec container_file(Client.t(), binary(), String.t()) :: {:ok, [String.t()]} | {:error, term()}
  def container_file(%Client{} = client, file_data, filename) do
    multipart = [
      {:file, file_data, filename: filename}
    ]

    case Client.put_multipart(client, "/unrestrict/containerFile", multipart) do
      {:ok, body} -> {:ok, body}
      {:error, reason} -> {:error, reason}
    end
  end

  @doc """
  Unrestrict a container file by link.

  ## Parameters

    - `client` - The Real-Debrid client
    - `link` - The link to the container file

  ## Returns

    - `{:ok, links}` - List of unrestricted links
    - `{:error, reason}` on failure
  """
  @spec container_link(Client.t(), String.t()) :: {:ok, [String.t()]} | {:error, term()}
  def container_link(%Client{} = client, link) do
    form = %{"link" => link}

    case Client.post(client, "/unrestrict/containerLink", form: form) do
      {:ok, body} -> {:ok, body}
      {:error, reason} -> {:error, reason}
    end
  end
end
