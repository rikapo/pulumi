// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.My8110
{
    [EnumType]
    public readonly struct MyEnum : IEquatable<MyEnum>
    {
        private readonly string _value;

        private MyEnum(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MyEnum One { get; } = new MyEnum("one");
        public static MyEnum Two { get; } = new MyEnum("two");

        public static bool operator ==(MyEnum left, MyEnum right) => left.Equals(right);
        public static bool operator !=(MyEnum left, MyEnum right) => !left.Equals(right);

        public static explicit operator string(MyEnum value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MyEnum other && Equals(other);
        public bool Equals(MyEnum other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}