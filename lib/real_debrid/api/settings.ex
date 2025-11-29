defmodule RealDebrid.Api.Settings do
  @moduledoc """
  Settings management API functions.
  """

  alias RealDebrid.Client

  @type settings :: %{
          download_ports: [String.t()],
          download_port: String.t(),
          locales: map(),
          locale: String.t(),
          streaming_qualities: [String.t()],
          streaming_quality: String.t(),
          mobile_streaming_quality: String.t(),
          streaming_languages: map(),
          streaming_language_preference: String.t(),
          streaming_cast_audio: [String.t()],
          streaming_cast_audio_preference: String.t()
        }

  @doc """
  Get current user settings.

  ## Returns

    - `{:ok, settings}` on success
    - `{:error, reason}` on failure
  """
  @spec get(Client.t()) :: {:ok, settings()} | {:error, term()}
  def get(%Client{} = client) do
    case Client.get(client, "/settings") do
      {:ok, body, _headers} ->
        {:ok,
         %{
           download_ports: body["download_ports"] || [],
           download_port: body["download_port"],
           locales: body["locales"] || %{},
           locale: body["locale"],
           streaming_qualities: body["streaming_qualities"] || [],
           streaming_quality: body["streaming_quality"],
           mobile_streaming_quality: body["mobile_streaming_quality"],
           streaming_languages: body["streaming_languages"] || %{},
           streaming_language_preference: body["streaming_language_preference"],
           streaming_cast_audio: body["streaming_cast_audio"] || [],
           streaming_cast_audio_preference: body["streaming_cast_audio_preference"]
         }}

      {:error, reason} ->
        {:error, reason}
    end
  end

  @doc """
  Update a setting.

  ## Parameters

    - `client` - The Real-Debrid client
    - `name` - Setting name
    - `value` - Setting value

  ## Returns

    - `:ok` on success
    - `{:error, reason}` on failure
  """
  @spec update(Client.t(), String.t(), String.t()) :: :ok | {:error, term()}
  def update(%Client{} = client, name, value) do
    form = %{"setting_name" => name, "setting_value" => value}

    case Client.post(client, "/settings/update", form: form, expected_status: 204) do
      {:ok, _body} -> :ok
      {:error, reason} -> {:error, reason}
    end
  end

  @doc """
  Convert fidelity points to premium days.

  ## Returns

    - `:ok` on success
    - `{:error, reason}` on failure
  """
  @spec convert_points(Client.t()) :: :ok | {:error, term()}
  def convert_points(%Client{} = client) do
    case Client.post(client, "/settings/convertPoints", expected_status: 204) do
      {:ok, _body} -> :ok
      {:error, reason} -> {:error, reason}
    end
  end

  @doc """
  Send a password reset email.

  ## Returns

    - `:ok` on success
    - `{:error, reason}` on failure
  """
  @spec change_password(Client.t()) :: :ok | {:error, term()}
  def change_password(%Client{} = client) do
    case Client.post(client, "/settings/changePassword", expected_status: 204) do
      {:ok, _body} -> :ok
      {:error, reason} -> {:error, reason}
    end
  end
end
