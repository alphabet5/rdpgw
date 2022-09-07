package protocol

const (
	PKT_TYPE_HANDSHAKE_REQUEST      = 0x1
	PKT_TYPE_HANDSHAKE_RESPONSE     = 0x2
	PKT_TYPE_EXTENDED_AUTH_MSG      = 0x3
	PKT_TYPE_TUNNEL_CREATE          = 0x4
	PKT_TYPE_TUNNEL_RESPONSE        = 0x5
	PKT_TYPE_TUNNEL_AUTH            = 0x6
	PKT_TYPE_TUNNEL_AUTH_RESPONSE   = 0x7
	PKT_TYPE_CHANNEL_CREATE         = 0x8
	PKT_TYPE_CHANNEL_RESPONSE       = 0x9
	PKT_TYPE_DATA                   = 0xA
	PKT_TYPE_SERVICE_MESSAGE        = 0xB
	PKT_TYPE_REAUTH_MESSAGE         = 0xC
	PKT_TYPE_KEEPALIVE              = 0xD
	PKT_TYPE_CLOSE_CHANNEL          = 0x10
	PKT_TYPE_CLOSE_CHANNEL_RESPONSE = 0x11
)

const (
	HTTP_TUNNEL_RESPONSE_FIELD_TUNNEL_ID   = 0x01
	HTTP_TUNNEL_RESPONSE_FIELD_CAPS        = 0x02
	HTTP_TUNNEL_RESPONSE_FIELD_SOH_REQ     = 0x04
	HTTP_TUNNEL_RESPONSE_FIELD_CONSENT_MSG = 0x10
)

const (
	HTTP_EXTENDED_AUTH_NONE      = 0x0
	HTTP_EXTENDED_AUTH_SC        = 0x1  /* Smart card authentication. */
	HTTP_EXTENDED_AUTH_PAA       = 0x02 /* Pluggable authentication. */
	HTTP_EXTENDED_AUTH_SSPI_NTLM = 0x04 /* NTLM extended authentication. */
)

const (
	HTTP_TUNNEL_AUTH_RESPONSE_FIELD_REDIR_FLAGS  = 0x01
	HTTP_TUNNEL_AUTH_RESPONSE_FIELD_IDLE_TIMEOUT = 0x02
	HTTP_TUNNEL_AUTH_RESPONSE_FIELD_SOH_RESPONSE = 0x04
)

const (
	HTTP_TUNNEL_REDIR_ENABLE_ALL        = 0x80000000
	HTTP_TUNNEL_REDIR_DISABLE_ALL       = 0x40000000
	HTTP_TUNNEL_REDIR_DISABLE_DRIVE     = 0x01
	HTTP_TUNNEL_REDIR_DISABLE_PRINTER   = 0x02
	HTTP_TUNNEL_REDIR_DISABLE_PORT      = 0x04
	HTTP_TUNNEL_REDIR_DISABLE_CLIPBOARD = 0x08
	HTTP_TUNNEL_REDIR_DISABLE_PNP       = 0x10
)

const (
	HTTP_CHANNEL_RESPONSE_FIELD_CHANNELID   = 0x01
	HTTP_CHANNEL_RESPONSE_FIELD_AUTHNCOOKIE = 0x02
	HTTP_CHANNEL_RESPONSE_FIELD_UDPPORT     = 0x04
)

const (
	HTTP_TUNNEL_PACKET_FIELD_PAA_COOKIE = 0x1
)

const (
	SERVER_STATE_INITIALIZED = 0x0
	SERVER_STATE_HANDSHAKE   = 0x1
	SERVER_STATE_TUNNEL_CREATE    = 0x2
	SERVER_STATE_TUNNEL_AUTHORIZE = 0x3
	SERVER_STATE_CHANNEL_CREATE   = 0x4
	SERVER_STATE_OPENED           = 0x5
	SERVER_STATE_CLOSED           = 0x6
)

const (
	HTTP_CAPABILITY_TYPE_QUAR_SOH          = 0x1
	HTTP_CAPABILITY_IDLE_TIMEOUT           = 0x2
	HTTP_CAPABILITY_MESSAGING_CONSENT_SIGN = 0x4
	HTTP_CAPABILITY_MESSAGING_SERVICE_MSG  = 0x8
	HTTP_CAPABILITY_REAUTH                 = 0x10
	HTTP_CAPABILITY_UDP_TRANSPORT          = 0x20
)
