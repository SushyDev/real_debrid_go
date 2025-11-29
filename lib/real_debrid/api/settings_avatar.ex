defmodule RealDebrid.Api.SettingsAvatar do
  @moduledoc """
  Avatar management API functions.
  """

  alias RealDebrid.Client

  @doc """
  Upload an avatar file.

  ## Parameters

    - `client` - The Real-Debrid client
    - `file_data` - Binary content of the image file
    - `filename` - Name of the file

  ## Returns

    - `:ok` on success
    - `{:error, reason}` on failure
  """
  @spec upload(Client.t(), binary(), String.t()) :: :ok | {:error, term()}
  def upload(%Client{} = client, file_data, filename) do
    multipart = [
      {:file, file_data, filename: filename, content_type: "image/*"}
    ]

    case Client.put_multipart(client, "/settings/avatarFile", multipart) do
      {:ok, _body} -> :ok
      {:error, reason} -> {:error, reason}
    end
  end

  @doc """
  Delete the current avatar.

  ## Returns

    - `:ok` on success
    - `{:error, reason}` on failure
  """
  @spec delete(Client.t()) :: :ok | {:error, term()}
  def delete(%Client{} = client) do
    Client.delete(client, "/settings/avatarDelete")
  end
end
